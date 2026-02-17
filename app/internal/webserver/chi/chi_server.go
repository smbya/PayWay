package chi

import (
	"context"
	"log"
	nethttp "net/http"
	"payway/internal/controller/http"

	"github.com/go-chi/chi/v5"
)

type ChiServer struct {
	ctx    context.Context
	port   string
	router *chi.Mux
}

func NewChiServer(ctx context.Context, port string) *ChiServer {
	return &ChiServer{
		ctx:    ctx,
		port:   port,
		router: chi.NewRouter(),
	}
}

func (s *ChiServer) RegisterRoutes(routes []http.Route) {
	for _, route := range routes {
		HandlerFunc := func(w nethttp.ResponseWriter, r *nethttp.Request) {
			log.Print("opa")
			params := make(map[string]string)
			rctx := chi.RouteContext(r.Context())
			if rctx != nil {
				for i, key := range rctx.URLParams.Keys {
					params[key] = rctx.URLParams.Values[i]
				}
			}

			actionCtx := http.HttpContext{
				UrlParams: params,
				Body:      "", //r.Body,
			}

			response, statusCode, _ := route.Handler(actionCtx)

			w.WriteHeader(statusCode)
			w.Write([]byte(response))
		}

		log.Print(route.Method,
			route.Path)

		s.router.MethodFunc(
			route.Method,
			route.Path,
			HandlerFunc,
		)
	}
}

func (s *ChiServer) Run() error {
	return nethttp.ListenAndServe(":"+s.port, s.router)
}
