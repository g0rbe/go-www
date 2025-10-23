package www_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication"
)

func ExampleNewClientWithAuthentication_bearerToken() {

	c := www.NewClientWithAuthentication(authentication.NewBearerToken("bearer-token"))

	// Issues an HTTP with Bearer Authentication
	c.Get("https://example.com/api/path")
}

func ExampleNewClientWithAuthentication_xApiKey() {

	c := www.NewClientWithAuthentication(authentication.NewXApiKey("api-key"))

	// Issues an HTTP with X-Api-Key set to "api-key"
	c.Get("https://example.com/api/path")
}
