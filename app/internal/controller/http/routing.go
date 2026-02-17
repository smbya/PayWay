package http

import "log"

func GetRoutes(facade PaymentFacade) []Route {
	return []Route{
		{
			Method: "GET",
			Path:   "/test/{name}/{name2}",
			Handler: func(c HttpContext) (string, int, error) {
				return "test with params: " + c.UrlParams["name"] + " and " + c.UrlParams["name2"], 200, nil
			},
		},
		{
			Method: "POST",
			Path:   "/payments",
			Handler: func(c HttpContext) (string, int, error) {
				log.Print("hello handler")
				return "Hello world", 200, nil
			},
		},
		{
			Method:  "GET",
			Path:    "/payments/{id}",
			Handler: facade.GetPaymentStatus,
		},
	}
}
