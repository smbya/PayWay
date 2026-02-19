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
			Method:  "POST",
			Path:    "/payments",
			Handler: facade.CreatePayment,
		},
		{
			Method: "POST",
			Path:   "/post",
			Handler: func(c HttpContext) (string, int, error) {
				log.Print("post test")
				return "post request", 200, nil
			},
		},
		{
			Method:  "GET",
			Path:    "/payments/{id}",
			Handler: facade.GetPaymentStatus,
		},
	}
}
