package www

import (
	"io"
	"net/http"
)

type Response struct {
	Status int         // Status code
	Header http.Header // Response headers
	Body   []byte      // Body of the response
}

// ParseResponse parses a *[http.Response].
//
// This function closes the [http.Response.Body].
func ParseResponse(resp *http.Response) (*Response, error) {

	defer resp.Body.Close()

	var (
		err error
		r   *Response = new(Response)
	)

	r.Status = resp.StatusCode
	r.Header = resp.Header

	r.Body, err = io.ReadAll(resp.Body)

	return r, err
}
