package authentication_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication"
)

func ExampleBearerToken() {

	bt := authentication.NewBearerToken("sup3rs3cr3t")

	client := www.NewClientWithAuthentication(bt)

	client.Get("https://example.com/api/path")
	// GET /api/path HTTP/3
	//
	// Host: example.com
	// Authorization: Bearer sup3rs3cr3t
	// ...
}
