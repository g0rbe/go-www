package www

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	BaseURL        string          // Base URL (eg.: "https://example.com")
	ContentType    string          // If set, sets the requests "Content-Type" header to the given value
	Authenticators []Authenticator // A slice of Authenticators to authenticate the requests
}

func NewClient(baseUrl string, contentType string, authenticators ...Authenticator) *Client {
	return &Client{
		BaseURL:        strings.TrimSuffix(baseUrl, "/"),
		ContentType:    contentType,
		Authenticators: authenticators,
	}
}

// NewRequest returns a new [http.Request] with Content-Type header and authentication headers set.
func (c *Client) NewRequest(method string, path string, body io.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, c.BaseURL+path, body)
	if err != nil {
		return nil, err
	}

	if len(c.ContentType) != 0 {
		req.Header.Set("Content-Type", c.ContentType)
	}

	for i := range c.Authenticators {
		c.Authenticators[i].SetHeader(req)
	}

	return req, nil
}

// Get do a HTTP GET request to the  BaseURL + path.
//
// Returns the response status code and the body bytes.
func (c *Client) Get(path string) (int, []byte, error) {

	req, err := c.NewRequest("GET", c.BaseURL+path, nil)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read body: %v", err)
	}

	return resp.StatusCode, buf, err
}

// Post do a HTTP POST request to the BaseURL + path.
//
// The JSON encoded v is set to request body.
//
// Returns the response status code and the body bytes.
func (c *Client) Post(path string, v any) (int, []byte, error) {

	buf, err := json.Marshal(v)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	req, err := c.NewRequest("POST", c.BaseURL+path, bytes.NewBuffer(buf))
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read body: %v", err)
	}

	return resp.StatusCode, buf, err
}

// Put do a HTTP PUT request to the BaseURL + path.
//
// The JSON encoded v is set to request body.
//
// Returns the response status code and the body bytes.
func (c *Client) Put(path string, v any) (int, []byte, error) {

	buf, err := json.Marshal(v)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	req, err := c.NewRequest("PUT", c.BaseURL+path, bytes.NewBuffer(buf))
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read body: %v", err)
	}

	return resp.StatusCode, body, nil
}
