package www

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client http.Client

// NewClient returns a new *[Client].
func NewClient() *Client {

	return (*Client)(&http.Client{})
}

// NewClientWithXApiKey returns a new *[Client] that
// authenticates every request by setting the "X-Api-Key" header's value to the specified API key.
func NewClientWithAuthentication(roundtrip http.RoundTripper) *Client {
	return (*Client)(&http.Client{Transport: roundtrip})
}

// Do sends an HTTP request and returns the respons status code, the response headers and the body bytes.
func (c *Client) Do(req *http.Request) (*Response, error) {

	resp, err := (*http.Client)(c).Do(req)
	if err != nil {
		return nil, err
	}

	return ParseResponse(resp)
}

// Get issues a HTTP GET request to the specified URL.
//
// Returns the response status code, the response headers and the body bytes.
func (c *Client) Get(url string) (*Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	return c.Do(req)
}

// Post issues a HTTP POST request to the specified URL.
//
// This function sets the "Content-Type" header to the specified value in contentType.
//
// Returns the response status code, the response headers and the body bytes.
func (c *Client) Post(url string, contentType string, body io.Reader) (*Response, error) {

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)

	return c.Do(req)
}

// PostJSON issues a HTTP POST request to the specified URL.
//
// The JSON encoded v is marshalled set as the request body.
// The "Content-Type" header in the request is set to [ContentTypeJSON].
//
// Returns the response status code, the response headers and the body bytes.
func (c *Client) PostJSON(url string, v any) (*Response, error) {

	buf, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return c.Post(url, ContentTypeJSON, bytes.NewBuffer(buf))
}

// Put issues a HTTP PUT request to the specified URL.
//
// This function sets the "Content-Type" header to the specified value in contentType.
//
// Returns the response status code, the response headers and the body bytes.
func (c *Client) Put(url string, contentType string, body io.Reader) (*Response, error) {

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)

	return c.Do(req)
}

// PutJSON issues a HTTP PUT request to the specified URL.
//
// The JSON encoded v is marshalled set as the request body.
// The "Content-Type" header in the request is set to [ContentTypeJSON].
//
// Returns the response status code, the response headers and the body bytes.
func (c *Client) PutJSON(url string, v any) (*Response, error) {

	buf, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return c.Put(url, ContentTypeJSON, bytes.NewBuffer(buf))
}
