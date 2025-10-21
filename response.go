package www

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"mime"
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

// Unmarshal unmarshals the [Response.Body] based on the response `Content-Type`:
//   - [json.Unmarhsal] if `Content-Type` is `application/json`
//   - [xml.Unmarshal]  if `Content-Type` is `application/xml`
//   - [error] otherwise
func (r *Response) Unmarshal(v any) error {

	contentType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return fmt.Errorf("failed to parse Content-Type: %w", err)
	}

	switch contentType {
	case ContentTypeJSON:
		return json.Unmarshal(r.Body, v)
	case ContentTypeXML:
		return xml.Unmarshal(r.Body, v)
	default:
		return fmt.Errorf("unknown Content-Type: %s", contentType)
	}
}
