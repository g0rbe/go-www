package bearertoken_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication/bearertoken"
)

func ExampleBearerToken() {

	bt := bearertoken.New("sup3rs3cr3t")

	client := www.NewClientWithAuthentication(bt)

	client.Get("https://example.com/api/path")
	// GET /api/path HTTP/3
	//
	// Host: example.com
	// Authorization: Bearer sup3rs3cr3t
	// ...
}
