package authentication_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication"
)

func ExampleXApiKey() {

	x := authentication.NewXApiKey("sup3rs3cr3t")

	client := www.NewClientWithAuthentication(x)

	client.Get("https://example.com/api/path")
	// GET /api/path HTTP/3
	//
	// Host: example.com
	// X-Api-Key: sup3rs3cr3t
	// ...
}
