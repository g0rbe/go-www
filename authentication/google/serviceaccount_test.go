package google_test

import (
	"git.gorbe.io/go/www"
	"git.gorbe.io/go/www/authentication/google"
)

func ExampleServiceAccount_serviceAccountFromJSON() {

	jsonKey := []byte(
		`
{
  "type": "...",
  "project_id": "...",
  "private_key_id": "",
  "private_key": "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n"
  "client_email": "...",
  "client_id": "...",
  "auth_uri": "...",
  "token_uri": "...",
  "auth_provider_x509_cert_url": "...",
  "client_x509_cert_url": "...",
  "universe_domain": "..."
}
`)

	client := www.NewClient()

	sa, err := google.ServiceAccountFromJSON(jsonKey)
	if err != nil {
		/// handle error
	}

	client.SetAuthentication(sa)

	client.Get("...")
}

func ExampleServiceAccount_serviceAccountFromJSONFile() {

	client := www.NewClient()

	sa, err := google.ServiceAccountFromJSONFile("/path/to/key.json")
	if err != nil {
		/// handle error
	}

	client.SetAuthentication(sa)

	client.Get("...")
}
