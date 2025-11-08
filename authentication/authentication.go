// Package authentication provides HTTP authentication
package authentication
import "net/http"

type Authentication = http.RoundTripper
