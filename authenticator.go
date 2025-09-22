package www

import "net/http"

type Authenticator interface {
	SetHeader(req *http.Request)
}

type AuthenticationHeader struct {
	Key   string
	Value string
}

func NewAuthenticationHeader(key string, value string) *AuthenticationHeader {
	return &AuthenticationHeader{Key: key, Value: value}
}

func (h *AuthenticationHeader) SetHeader(r *http.Request) {
	r.Header.Set(h.Key, h.Value)
}
