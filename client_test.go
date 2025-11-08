package www_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication/bearertoken"
	"git.gorbe.io/go/www/authentication/google"
	"git.gorbe.io/go/www/authentication/xapikey"
)

func ExampleNewClient_bearerToken() {

	c := www.NewClient()
	c.SetAuthentication(bearertoken.New("bearer-token"))

	// Issues an HTTP with Bearer Authentication
	c.Get("https://example.com/api/path")
}

func ExampleNewClient_xApiKey() {

	c := www.NewClient()
	c.SetAuthentication(xapikey.New("api-key"))

	// Issues an HTTP with X-Api-Key set to "api-key"
	c.Get("https://example.com/api/path")
}

func ExampleNewClient_googleServiceAccount() {

	c := www.NewClient()

	sa, err := google.ServiceAccountFromJSONFile("/path/to/key.json")
	if err != nil {
		/// handle error
	}

	c.SetAuthentication(sa)

	c.Get("https://example.com/api/path")
}
