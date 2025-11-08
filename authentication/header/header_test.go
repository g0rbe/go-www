package header_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication/header"
)

func ExampleHeader() {

	client := www.NewClient()

	client.SetAuthentication(header.New("Authentication-Header", "sup3rs3cr3t"))

	client.Get("https://example.com/api/path")
	// GET /api/path HTTP/3
	//
	// Host: example.com
	// Authentication-Header: sup3rs3cr3t
	// ...
}
