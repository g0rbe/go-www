package www_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication/bearertoken"
	"git.gorbe.io/go/www/authentication/google"
	"git.gorbe.io/go/www/authentication/xapikey"
)

func ExampleNewClientWithAuthentication_bearerToken() {

	c := www.NewClientWithAuthentication(bearertoken.New("bearer-token"))

	// Issues an HTTP with Bearer Authentication
	c.Get("https://example.com/api/path")
}

func ExampleNewClientWithAuthentication_xApiKey() {

	c := www.NewClientWithAuthentication(xapikey.New("api-key"))

	// Issues an HTTP with X-Api-Key set to "api-key"
	c.Get("https://example.com/api/path")
}

func ExampleNewClientWithAuthentication_googleServiceAccount() {

	sa, err := google.ServiceAccountFromJSONFile("/path/to/key.json")
	if err != nil {
		/// handle error
	}

	c := www.NewClientWithAuthentication(sa)

	c.Get("https://example.com/api/path")
}
