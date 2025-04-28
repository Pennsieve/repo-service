package response

const HeaderAuthorization = "Authorization"
const HeaderContentType = "Content-Type"
const HeaderAccept = "Accept"
const HeaderContentEncoding = "Content-Encoding"

type Headers map[string]string

var NoHeaders = Headers{}

func StandardResponseHeaders(custom Headers) map[string]string {
	headers := make(Headers)
	// copy in any custom headers
	for key, value := range custom {
		headers[key] = value
	}
	// and add the set of standard response headers
	headers[HeaderContentType] = "application/json; charset=utf-8"
	return headers
}
