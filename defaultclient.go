package www

import (
	"io"
	"net/http"
	"net/url"
)

var (
	DefaultClient = NewClient()
)

// Do sends an HTTP request with [DefaultClient].
func Do(req *http.Request) (*Response, error) {
	return DefaultClient.Do(req)
}

// Get uses [DefaultClient] to issues a HTTP GET request to the specified URL.
func Get(url string) (*Response, error) {
	return DefaultClient.Get(url)
}

// Post uses [DefaultClient] to issues a HTTP POST request to the specified URL.
//
// This function sets the "Content-Type" header to the specified value in contentType.
func Post(url string, contentType string, body io.Reader) (*Response, error) {
	return DefaultClient.Post(url, contentType, body)
}

// PostJSON uses [DefaultClient] to issues a HTTP POST request to the specified URL,
// with v JSON encoded as the request body.
//
// The "Content-Type" header in the request is set to [ContentTypeJSON].
func PostJSON(url string, v any) (*Response, error) {
	return DefaultClient.PostJSON(url, v)
}

// Put uses [DefaultClient] to issues a HTTP PUT request to the specified URL.
//
// This function sets the "Content-Type" header to the specified value in contentType.
func Put(url string, contentType string, body io.Reader) (*Response, error) {
	return DefaultClient.Put(url, contentType, body)
}

// PutJSON uses [DefaultClient] to issues a HTTP PUT request to the specified URL.,
// with v JSON encoded as the request body.
//
// The "Content-Type" header in the request is set to [ContentTypeJSON].
func PutJSON(url string, v any) (*Response, error) {
	return DefaultClient.PutJSON(url, v)
}

// PostForm uses [DefaultClient] to issues a POST to the specified URL,
// with data's keys and values URL-encoded as the request body.
//
// The Content-Type header is set to `application/x-www-form-urlencoded`.
func PostForm(url string, data url.Values) (*Response, error) {
	return DefaultClient.PostForm(url, data)
}
