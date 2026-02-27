package webserver

import (
	"context"
	"log/slog"

	"payway/internal/controller/http"
	"payway/internal/controller/webserver/chi"
	"payway/internal/controller/webserver/gin"
)

type WebServer interface {
	Serve(ctx context.Context) error
}

func NewWebServer(
	handlerType string,
	port string,
	logger *slog.Logger,
	facade http.PaymentFacade,
) WebServer {
	switch handlerType {
	case "chi":
		return chi.NewChiServer(port, logger, facade)
	case "gin":
		return gin.NewGinServer(port, logger, facade)
	default:
		logger.Error("Server not choosed: ", "handlerType", handlerType)
		panic("Server not choosed: " + handlerType)
	}
}
