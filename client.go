package censys

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	rootURL = "https://search.censys.io/api"
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

	hostResp := &Host{}
	err = json.Unmarshal(*jsonResp.Result, hostResp)
	if err != nil {
		return nil, err
	}

	return hostResp, nil
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
