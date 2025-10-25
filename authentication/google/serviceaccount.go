package google

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
)

// ServiceAccount uses a Google Developers service account JSON key to read the credentials that authorize and authenticate the requests.
type ServiceAccount struct {
	Email        string   `json:"client_email" yaml:"Email"`
	PrivateKey   string   `json:"private_key" yaml:"PrivateKey"`
	PrivateKeyID string   `json:"private_key_id" yaml:"PrivateKeyID"`
	TokenURL     string   `json:"token_uri" yaml:"TokenURL"`
	Scopes       []string `json:"scopes" yaml:"Scopes"`
	Subject      string   `json:"subject" yaml:"Subject"`

	// Transport is the underlying HTTP transport to use when making requests.
	// Defaults to [http.DefaultTransport]
	Transport http.RoundTripper
}

// ServiceAccountFromJSON uses a Google Developers service account JSON key to read the credentials that authorize and authenticate the requests.
func ServiceAccountFromJSON(data []byte) (*ServiceAccount, error) {

	sa := new(ServiceAccount)

	err := json.Unmarshal(data, sa)
	if err != nil {
		return nil, err
	}

	return sa, nil
}

// ServiceAccountFromJSONFile uses a Google Developers service account JSON key file in path name to read the credentials that authorize and authenticate the requests.
func ServiceAccountFromJSONFile(name string) (*ServiceAccount, error) {

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", name, err)
	}

	return ServiceAccountFromJSON(data)
}

// transport returns the [ServiceAccount.Transport] if set
// or
// [http.DefaultTransport] if nil.
func (sa *ServiceAccount) transport() http.RoundTripper {
	if sa.Transport != nil {
		return sa.Transport
	}

	return http.DefaultTransport
}

// Token implements [golang.org/x/oauth2.TokenSource] interface.
func (sa *ServiceAccount) Token() (*oauth2.Token, error) {

	jwtConf := &jwt.Config{
		Email:        sa.Email,
		PrivateKey:   []byte(sa.PrivateKey),
		PrivateKeyID: sa.PrivateKeyID,
		TokenURL:     sa.TokenURL,
		Scopes:       sa.Scopes,
		Subject:      sa.Subject,
	}

	return jwtConf.TokenSource(context.TODO()).Token()
}

// RoundTrip implements [http.RoundTripper] interface.
func (sa *ServiceAccount) RoundTrip(req *http.Request) (*http.Response, error) {

	token, err := sa.Token()
	if err != nil {
		return nil, err
	}

	req2 := req.Clone(req.Context())

	token.SetAuthHeader(req2)

	return sa.transport().RoundTrip(req2)
}
