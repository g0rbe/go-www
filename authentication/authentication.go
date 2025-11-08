/*
Package authentication provides HTTP authentication

The authentications are implemented with the [http.RoundTripper].
The authentication header are added to every request with the RoundTrip() method.
*/
package authentication

import "net/http"

// Authentication is an alias type to [http.RoundTripper].
type Authentication = http.RoundTripper
