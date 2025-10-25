package header_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication/header"
)

func ExampleHeader() {

	h := header.New("Authentication-Header", "sup3rs3cr3t")

	client := www.NewClientWithAuthentication(h)

	client.Get("https://example.com/api/path")
	// GET /api/path HTTP/3
	//
	// Host: example.com
	// Authentication-Header: sup3rs3cr3t
	// ...
}
