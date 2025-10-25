package xapikey_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication/xapikey"
)

func ExampleXApiKey() {

	x := xapikey.New("sup3rs3cr3t")

	client := www.NewClientWithAuthentication(x)

	client.Get("https://example.com/api/path")
	// GET /api/path HTTP/3
	//
	// Host: example.com
	// X-Api-Key: sup3rs3cr3t
	// ...
}
