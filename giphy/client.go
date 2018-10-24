package giphy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/danew/giffy/util"
)

// DefaultClient is the default Giphy API client
var DefaultClient = NewClient()

var PublicBetaAPIKey = "dc6zaTOxFJmzC"

// A Client communicates with the Giphy API.
type Client struct {
	apiKey     string
	limit      int
	rating     string
	baseURL    *url.URL
	basePath   string
	userAgent  string
	httpClient *http.Client
}

// NewClient returns a new Giphy API client.
func NewClient() *Client {
	var httpClient *http.Client
	cloned := *http.DefaultClient
	httpClient = &cloned

	c := &Client{
		apiKey: util.Env("GIPHY_API_KEY", PublicBetaAPIKey),
		rating: util.Env("GIPHY_RATING", "g"),
		limit:  util.EnvInt("GIPHY_LIMIT", 25),
		baseURL: &url.URL{
			Scheme: util.Env("GIPHY_BASE_URL_SCHEME", "https"),
			Host:   util.Env("GIPHY_BASE_URL_HOST", "api.giphy.com"),
		},
		basePath:   util.Env("GIPHY_BASE_PATH", "/v1"),
		userAgent:  util.Env("GIPHY_USER_AGENT", "giffy"),
		httpClient: httpClient,
	}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(s string) (*http.Request, error) {
	rel, err := url.Parse(c.basePath + s)
	if err != nil {
		return nil, err
	}

	q := rel.Query()
	q.Set("api_key", c.apiKey)
	q.Set("rating", c.rating)
	rel.RawQuery = q.Encode()

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.userAgent)
	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	req.Close = true
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	if err != nil {
		return nil, fmt.Errorf("Error reading response from %s %s: %s", req.Method, req.URL.RequestURI(), err)
	}
	return resp, nil
}
