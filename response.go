package goResponse

type Response struct {
	headers    map[string]string
	body       string
	statusCode int
}

func NewGoResponse() *Response {
	return &Response{
		headers: map[string]string{},
	}
}

func (r *Response) GetHeaders() map[string]string {
	return r.headers
}

func (r *Response) HasHeaders(name string) bool {
	_, ok := r.headers[name]

	return ok
}

func (r *Response) GetHeader(name string) string {
	header, ok := r.headers[name]

	if !ok {
		header = ""
	}

	return header
}

func (r *Response) WithHeader(name string, value string) {
	r.headers[name] = value
}

func (r *Response) WithBody(value string) {
	r.body = value
}

func (r *Response) GetBody() string {
	return r.body
}

func (r *Response) GetStatusCode() int {
	return r.statusCode
}

func (r *Response) WithStatusCode(code int) {
	r.statusCode = code
}
