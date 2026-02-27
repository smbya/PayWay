package chi

import (
	"context"
	"log"

	// "log"
	"log/slog"
	nethttp "net/http"
	"payway/internal/controller/http"

	"github.com/go-chi/chi/v5"
)

type ChiServer struct {
	ctx    context.Context
	port   string
	router *chi.Mux
	logger *slog.Logger
}

func NewChiServer(ctx context.Context, port string, logger *slog.Logger) *ChiServer {
	return &ChiServer{
		ctx:    ctx,
		port:   port,
		router: chi.NewRouter(),
		logger: logger,
	}
}

func (s *ChiServer) RegisterRoutes(routes []http.Route) {
	for _, route := range routes {
		HandlerFunc := func(w nethttp.ResponseWriter, r *nethttp.Request) {
			s.logger.Info("HTTP Request", "URL", r.URL.String(), "method", r.Method)
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

		s.logger.Info("Added new route", "method", route.Method, "path", route.Path)

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
