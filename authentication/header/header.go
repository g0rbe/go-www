// Package header provides HTTP authentication with arbitrary header key and value
package header

import "net/http"

// Header is an http.RoundTripper that authenticates all requests with an arbitrary header key and its value.
//
//	<key>: <value>
type Header struct {
	Key   string // The key of the authentication header (eg.: X-API-KEY)
	Value string // The value of the authentication header

	// Transport is the underlying HTTP transport to use when making requests.
	// Defaults to [http.DefaultTransport]
	Transport http.RoundTripper
}

// NewHeader returns a new *[Header].
func New(key string, value string) *Header {
	return &Header{Key: key, Value: value}
}

// transport returns the [Header.Transport] if set
// or
// [http.DefaultTransport] if nil.
func (h *Header) transport() http.RoundTripper {
	if h.Transport != nil {
		return h.Transport
	}

	return http.DefaultTransport
}

// RoundTrip implements the [http.RoundTripper] interface.
//
// To prevent modifying [http.Request], RoundTrip use [http.Client.Clone] to copy req
// and sets the [http.Header] for the cloned [http.Request].
func (h *Header) RoundTrip(req *http.Request) (*http.Response, error) {

	// Clone the request to not modify the original request
	req2 := req.Clone(req.Context())

	req2.Header.Set(h.Key, h.Value)

	return h.transport().RoundTrip(req2)
}
