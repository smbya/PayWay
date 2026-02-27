package chi

import (
	"context"
	"encoding/json"
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

func NewChiServer(port string, logger *slog.Logger, handler *httphandler.PaymentHandler) *ChiServer {
	s := &ChiServer{
		port:   port,
		router: chi.NewRouter(),
		logger: logger,
	}
	s.registerRoutes(handler)
	return s
}

func (s *ChiServer) registerRoutes(handler *httphandler.PaymentHandler) {
	s.router.Post("/payments", func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("HTTP Request", "URL", r.URL.String(), "method", r.Method)

		var req httphandler.CreatePaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		payment, err := handler.CreatePayment(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(payment)
	})

	s.router.Get("/payments/{id}", func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("HTTP Request", "URL", r.URL.String(), "method", r.Method)
		id := chi.URLParam(r, "id")

		payment, err := handler.GetPaymentByID(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payment)
	})
}

func (s *ChiServer) Serve(ctx context.Context) error {
	return http.ListenAndServe(":"+s.port, s.router)
}
