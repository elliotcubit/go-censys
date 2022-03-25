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

func (c *clientImpl) GetHost(ip string, name string, at *time.Time) (*Host, error) {
	url := rootURL + "/v2/hosts/" + ip
	if name != "" {
		url += "+" + name
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if at != nil {
		q := req.URL.Query()
		q.Add("at_time", (*at).Format(time.RFC3339))
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.authedRequest(req)
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

	hostResp := &Host{}
	err = json.Unmarshal(*jsonResp.Result, hostResp)
	if err != nil {
		return nil, err
	}

	return hostResp, nil
}
