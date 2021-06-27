package goResponse

type Response struct {

}

func (r *Response) GetHeaders() map[string]string  {

}

func (r *Response) HasHeaders(name string) bool {

}

func (r *Response) GetHeader(name string) string  {

}

func (r *Response) WithHeader(name string, value string) {

}

func (r *Response) WithBody(value string)  {

}

func (r *Response) GetBody() string {

}

func (r *Response) GetStatusCode() int {

}

func (r *Response) WithStatusCode(code int) {

}