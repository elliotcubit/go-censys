package censys

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	rootURL = "https://search.censys.io/api"
)

const (
	VirtualHostsExclude = "EXCLUDE"
	VirtualHostsInclude = "INCLUDE"
	VirtualHostsOnly    = "ONLY"
)

const (
	SortRelevance  = "RELEVANCE"
	SortAscending  = "ASCENDING"
	SortDescending = "DESCENDING"
	SortRandom     = "RANDOM"
)

// NewClient creates a new Client with the provided API ID and Secret.
// If httpClient is provided, it will be used for all requests.
// Otherwise, a default is used.
func NewClient(id, secret string, httpClient *http.Client) Client {

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &clientImpl{
		clientID:     id,
		clientSecret: secret,
		httpClient:   httpClient,
	}
}

// Client is
type Client interface {
	GetHost(ip string, name string, at *time.Time) (*Host, error)

	// GetHostsMetadata fetches metadata about the hosts dataset,
	// for now, just the list of services that are being scanned.
	GetHostsMetadata() (*HostMetadata, error)

	// GetHostAggregate breaks the results of a query down by
	// the given field, and returns bucket counts for each potential
	// field value and the aggregate sum of results for the query.
	GetHostAggregate(
		query string,
		field string,
		numBuckets int,
		virtualHosts string,
	) (*HostAggregate, error)

	// SearchHosts fetches a page of search results for the given query.
	//
	// virtualHosts should be one of:
	//  - The empty string
	//  - VirtualHostsExclude
	//  - VirtualHostsInclude
	//  - VirtualHostsOnly
	//
	// When empty, VirtualHostsExclude will be used.
	//
	// Cursor may be empty or a valid cursor from a previous call to SearchHosts.
	SearchHosts(
		query string,
		perPage int,
		virtualHosts string,
		sortOrder string,
		cursor string,
	) (*SearchHostsResult, error)

	GetHostNames(
		ip string,
		perPage int,
		cursor string,
	) ([]string, string, error)
}

type clientImpl struct {
	clientID     string
	clientSecret string
	httpClient   *http.Client
}

func (c *clientImpl) authedRequest(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(c.clientID, c.clientSecret)
	return c.httpClient.Do(req)
}

func (c *clientImpl) getReq(path string, qargs map[string]string) (*response, error) {
	req, err := http.NewRequest(http.MethodGet, rootURL+path, nil)
	if err != nil {
		return nil, err
	}

	if len(qargs) > 0 {
		q := req.URL.Query()
		for k, v := range qargs {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.authedRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonResp := &response{}
	err = json.Unmarshal(bodyBytes, jsonResp)
	if err != nil {
		return nil, err
	}

	err = jsonResp.asError()
	if err != nil {
		return nil, err
	}

	return jsonResp, err
}

func (c *clientImpl) GetHost(ip string, name string, at *time.Time) (*Host, error) {
	path := "/v2/hosts/" + ip
	if name != "" {
		path += "+" + name
	}

	qargs := make(map[string]string)
	if at != nil {
		qargs["at_time"] = (*at).Format(time.RFC3339)
	}

	jsonResp, err := c.getReq(path, qargs)
	if err != nil {
		return nil, err
	}

	hostResp := &Host{}
	err = json.Unmarshal(*jsonResp.Result, hostResp)
	if err != nil {
		return nil, err
	}

	return hostResp, nil
}

func (c *clientImpl) GetHostAggregate(
	query string,
	field string,
	numBuckets int,
	virtualHosts string,
) (*HostAggregate, error) {
	path := "/v2/hosts/aggregate"

	if numBuckets <= 0 {
		numBuckets = 50
	}

	if virtualHosts == "" {
		virtualHosts = VirtualHostsExclude
	}

	qargs := map[string]string{
		"q":             query,
		"field":         field,
		"num_buckets":   strconv.Itoa(numBuckets),
		"virtual_hosts": virtualHosts,
	}

	jsonResp, err := c.getReq(path, qargs)
	if err != nil {
		return nil, err
	}

	retv := &HostAggregate{}
	err = json.Unmarshal(*jsonResp.Result, retv)
	if err != nil {
		return nil, err
	}

	return retv, nil
}

func (c *clientImpl) GetHostsMetadata() (*HostMetadata, error) {
	path := "/v2/metadata/hosts"
	jsonResp, err := c.getReq(path, nil)
	if err != nil {
		return nil, err
	}
	retv := &HostMetadata{}
	err = json.Unmarshal(*jsonResp.Result, retv)
	if err != nil {
		return nil, err
	}
	return retv, nil
}

func (c *clientImpl) SearchHosts(
	query string,
	perPage int,
	virtualHosts string,
	sortOrder string,
	cursor string,
) (*SearchHostsResult, error) {
	path := "/v2/hosts/search"
	if sortOrder == "" {
		sortOrder = SortRelevance
	}
	if virtualHosts == "" {
		virtualHosts = VirtualHostsExclude
	}
	if perPage > 100 {
		perPage = 100
	}
	if perPage <= 0 {
		perPage = 1
	}

	qargs := map[string]string{
		"q":             query,
		"per_page":      strconv.Itoa(perPage),
		"virtual_hosts": virtualHosts,
		"sort":          sortOrder,
	}

	if cursor != "" {
		qargs["cursor"] = cursor
	}

	jsonResp, err := c.getReq(path, qargs)
	if err != nil {
		return nil, err
	}

	retv := &SearchHostsResult{}
	err = json.Unmarshal(*jsonResp.Result, retv)
	if err != nil {
		return nil, err
	}

	return retv, nil
}

func (c *clientImpl) GetHostNames(
	ip string,
	perPage int,
	cursor string,
) (names []string, nextCursor string, err error) {
	path := "/v2/hosts/" + ip + "/names"
	if perPage > 100 {
		perPage = 100
	}
	if perPage <= 0 {
		perPage = 1
	}

	qargs := map[string]string{
		"per_page": strconv.Itoa(perPage),
	}

	if cursor != "" {
		qargs["cursor"] = cursor
	}

	jsonResp, err := c.getReq(path, qargs)
	if err != nil {
		return nil, "", err
	}

	parseRes := &struct {
		Names []string `json:"names"`
		Links *struct {
			Next string `json:"next"`
		} `json:"links"`
	}{}

	err = json.Unmarshal(*jsonResp.Result, parseRes)
	if err != nil {
		return nil, "", err
	}

	cursor = ""
	if parseRes.Links != nil {
		cursor = parseRes.Links.Next
	}

	return parseRes.Names, cursor, nil
}
