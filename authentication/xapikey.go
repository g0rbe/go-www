package authentication

import "net/http"

// XApiKey use `X-Api-Key` with the provided key to authenticate requests.
//
//	X-Api-Key: <key>
type XApiKey struct {
	Key string // The API key

	// Transport is the underlying HTTP transport to use when making requests.
	// Defaults to [http.DefaultTransport]
	Transport http.RoundTripper
}

// NewXApiKey returns a new *[XApiKey].
func NewXApiKey(key string) *XApiKey {
	return &XApiKey{Key: key}
}

// transport returns the [Header.Transport] if set
// or
// [http.DefaultTransport] if nil.
func (x *XApiKey) transport() http.RoundTripper {
	if x.Transport != nil {
		return x.Transport
	}

	return http.DefaultTransport
}

// RoundTrip implements the [http.RoundTripper] interface.
//
// To prevent modifying [http.Request], RoundTrip use [http.Client.Clone] to copy req
// and sets the [http.Header] for the cloned [http.Request].
func (x *XApiKey) RoundTrip(req *http.Request) (*http.Response, error) {

	// Clone the request to not modify the original request
	req2 := req.Clone(req.Context())

	req2.Header.Set("X-Api-Key", x.Key)

	return x.transport().RoundTrip(req2)
}
