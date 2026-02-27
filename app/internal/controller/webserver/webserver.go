package webserver

import (
	"context"
	"log/slog"
	"payway/internal/controller/http"
	"payway/internal/controller/webserver/chi"
	"payway/internal/controller/webserver/gin"
)

type WebServer interface {
	RegisterRoutes([]http.Route)
	Run() error
}

func NewWebServer(
	ctx context.Context,
	handlerType string,
	port string,
	logger *slog.Logger,
) WebServer {
	switch handlerType {
	case "chi":
		return chi.NewChiServer(ctx, port, logger)
	case "gin":
		return gin.NewGinServer(ctx, port)
	// TODO:
	// case "fasthttp":
	// 	return fastHttp.NewFastHttpServer(ctx, port)
	default:
		panic("Server not choosed: " + handlerType)
	}
}
