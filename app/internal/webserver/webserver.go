package webserver

import (
	"context"
	"payway/internal/controller/http"
	"payway/internal/webserver/chi"
	"payway/internal/webserver/gin"
)

type WebServer interface {
	RegisterRoutes([]http.Route)
	Run() error
}

func NewWebServer(
	ctx context.Context,
	handlerType string,
	port string,
) WebServer {
	switch handlerType {
	case "chi":
		return chi.NewChiServer(ctx, port)
	case "gin":
		return gin.NewGinServer(ctx, port)
	// TODO:
	// case "fasthttp":
	// 	return fastHttp.NewFastHttpServer(ctx, port)
	default:
		panic("Server not choosed: " + handlerType)
	}
}
