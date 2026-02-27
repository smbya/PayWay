package http

type HttpContext struct {
	UrlParams map[string]string
	Body      string
}
