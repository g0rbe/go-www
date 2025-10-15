package www

import "net/http"

// AuthenticationHeader is an http.RoundTripper that authenticates all requests
// with an arbitrary header key and its value.
type AuthenticationHeader struct {
	Key   string // The key of the authentication header (eg.: X-API-KEY)
	Value string // The value of the authentication header

	// Transport is the underlying HTTP transport to use when making requests.
	// Defaults to [http.DefaultTransport]
	Transport http.RoundTripper
}

// transport returns the [AuthenticationHeader.Transport] if set
// or
// [http.DefaultTransport] if nil.
func (ah *AuthenticationHeader) transport() http.RoundTripper {
	if ah.Transport != nil {
		return ah.Transport
	}

	return http.DefaultTransport
}

// RoundTrip implements the [http.RoundTripper] interface.
func (ah *AuthenticationHeader) RoundTrip(req *http.Request) (*http.Response, error) {

	// Clone the request to not modify the original request
	req2 := req.Clone(req.Context())

	req2.Header.Set(ah.Key, ah.Value)

	return ah.transport().RoundTrip(req2)
}

// Client returns an *http.Client that makes requests that are
// authenticated with an arbitrary key/value in the header.
func (ah *AuthenticationHeader) Client() *http.Client {
	return &http.Client{Transport: ah}
}
