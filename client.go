package www

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"git.gorbe.io/go/www/authentication"
)

var (
	ErrTooManyRedirects = errors.New("too many redirects")
)

type Client http.Client

// NewClient returns a new *[Client].
func NewClient() *Client {

	return (*Client)(&http.Client{})
}

// SetAuthentication sets *[Client] to authenticates every request with the given [git.gorbe.io/go/www/authentication.Authentication] auth.
func (c *Client) SetAuthentication(auth authentication.Authentication) {
	c.Transport = auth
}

// MaxRedirections sets the maximum number of redirections to follow.
//
// If maximum redirections reached, the Client's Get method returns
// both the Response (with its Body closed)
// and CheckRedirect's error (wrapped in a url.Error)
// instead of issuing the next req.
//
// This function panics if n is less than zero (n < 0)!
func (c *Client) MaxRedirections(n int) {
	if n < 0 {
		panic(fmt.Sprintf("invalid number for MaxRedirections: %d", n))
	}

	// Based on: https://cs.opensource.google/go/go/+/refs/tags/go1.25.4:src/net/http/client.go;l=820
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= n {
			return ErrTooManyRedirects
		}
		return nil
	}
}

// Do sends an HTTP request and returns *[Response].
func (c *Client) Do(req *http.Request) (*Response, error) {

	resp, err := (*http.Client)(c).Do(req)
	if err != nil {
		return nil, err
	}

	return ParseResponse(resp)
}

// Get issues a HTTP GET request to the specified URL.
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
func (c *Client) Post(url string, contentType string, body io.Reader) (*Response, error) {

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)

	return c.Do(req)
}

// PostJSON issues a HTTP POST request to the specified URL,
// with v JSON encoded as the request body.
//
// The "Content-Type" header in the request is set to [ContentTypeJSON].
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
func (c *Client) Put(url string, contentType string, body io.Reader) (*Response, error) {

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)

	return c.Do(req)
}

// PutJSON issues a HTTP PUT request to the specified URL,
// with v JSON encoded as the request body.
//
// The "Content-Type" header in the request is set to [ContentTypeJSON].
func (c *Client) PutJSON(url string, v any) (*Response, error) {

	buf, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return c.Put(url, ContentTypeJSON, bytes.NewBuffer(buf))
}

// PostForm issues a POST to the specified URL, with data's keys and
// values URL-encoded as the request body.
//
// The Content-Type header is set to `application/x-www-form-urlencoded`.
func (c *Client) PostForm(url string, data url.Values) (*Response, error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}
