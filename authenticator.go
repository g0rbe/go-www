package www

import "net/http"

type Authenticator interface {
	SetHeader(req *http.Request)
}

type AuthenticationHeader struct {
	Key   string
	Value string
}

// NewAuthenticationHeader returns a new AuthenticationHeader which implements the [git.gorbe.io/go/www.Authenticator] interface,
// thus every request has a Header key: value set.
func NewAuthenticationHeader(key string, value string) *AuthenticationHeader {
	return &AuthenticationHeader{Key: key, Value: value}
}

// SetHeader implements the [git.gorbe.io/go/www.Authenticator] interface.
func (h *AuthenticationHeader) SetHeader(r *http.Request) {
	r.Header.Set(h.Key, h.Value)
}
