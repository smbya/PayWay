package http

type HttpContext struct {
	//TODO: добавить сюда тоже контекст чтобы прокидывать в хэндлеры фасада
	UrlParams map[string]string
	Body      string
}

type HandlerFunc func(c HttpContext) (string, int, error)

type Route struct {
	Method  string
	Path    string
	Handler HandlerFunc
}
