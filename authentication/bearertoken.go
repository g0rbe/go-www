package authentication

import "net/http"

// BearerToken is an [http.RoundTripper] that use bearer authentication
// to authenticate every HTTP request.
//
//	Authorization: Bearer <token>
type BearerToken struct {
	Token string // The token to use in bearer authentication

	// Transport is the underlying HTTP transport to use when making requests.
	// Defaults to [http.DefaultTransport]
	Transport http.RoundTripper
}

// NewBearerToken use token in bearer authentication.
func NewBearerToken(token string) *BearerToken {
	return &BearerToken{Token: token}
}

// transport returns the [BearerToken.Transport] if set
// or
// [http.DefaultTransport] if nil.
func (bt *BearerToken) transport() http.RoundTripper {
	if bt.Transport != nil {
		return bt.Transport
	}

	return http.DefaultTransport
}

// RoundTrip implements the [http.RoundTripper] interface.
//
// To prevent modifying [http.Request], RoundTrip use [http.Client.Clone] to copy req
// and sets the [http.Header] for the cloned [http.Request].
func (bt *BearerToken) RoundTrip(req *http.Request) (*http.Response, error) {

	// Clone the request to not modify the original request
	req2 := req.Clone(req.Context())

	req2.Header.Set("Authorization", "Bearer "+bt.Token)

	return bt.transport().RoundTrip(req2)
}
