package www_test

import "git.gorbe.io/go/www"

func ExampleAuthenticationHeader() {

	ah := www.AuthenticationHeader{Key: "API-KEY", Value: "sup3rs3cr3t"}

	client := ah.Client()

	client.Get("https://example.com/api/path")
}
