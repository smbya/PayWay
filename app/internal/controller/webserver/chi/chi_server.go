package chi

import (
	"context"
	"log/slog"
	"net/http"

	httphandler "payway/internal/controller/http"

	"github.com/go-chi/chi/v5"
)

type ChiServer struct {
	port   string
	router *chi.Mux
	logger *slog.Logger
}

func NewChiServer(port string, logger *slog.Logger, facade httphandler.PaymentFacade) *ChiServer {
	s := &ChiServer{
		port:   port,
		router: chi.NewRouter(),
		logger: logger,
	}
	s.registerRoutes(facade)
	return s
}

func (s *ChiServer) registerRoutes(facade httphandler.PaymentFacade) {
	s.router.Get("/test/{name}/{name2}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		name2 := chi.URLParam(r, "name2")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test with params: " + name + " and " + name2))
	})

	s.router.Post("/payments", func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("HTTP Request", "URL", r.URL.String(), "method", r.Method)
		params := make(map[string]string)
		response, statusCode, _ := facade.CreatePayment(httphandler.HttpContext{
			UrlParams: params,
			Body:      "",
		})
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	})

	s.router.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("post request"))
	})

	s.router.Get("/payments/{id}", func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("HTTP Request", "URL", r.URL.String(), "method", r.Method)
		id := chi.URLParam(r, "id")
		response, statusCode, _ := facade.GetPaymentStatus(httphandler.HttpContext{
			UrlParams: map[string]string{"id": id},
			Body:      "",
		})
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	})
}

func (s *ChiServer) Serve(ctx context.Context) error {
	return http.ListenAndServe(":"+s.port, s.router)
}
