package goResponse

type CookieInterface interface {
	Set(r *ResponseInterface)
	Get(r *ResponseInterface, name string) (bool, *CookieInterface)
}

type ResponseInterface interface {
	GetHeaders() map[string]string
	HasHeaders(name string) bool
	GetHeader(name string) string
	WithHeader(name string, value string)
	WithBody(value string)
	GetBody() string
	GetStatusCode() int
	WithStatusCode(status int)
}