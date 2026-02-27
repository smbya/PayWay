package gin

import (
	"context"
	"log/slog"
	"net/http"

	httphandler "payway/internal/controller/http"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	port   string
	engine *gin.Engine
	logger *slog.Logger
}

func NewGinServer(port string, logger *slog.Logger, facade httphandler.PaymentFacade) *GinServer {
	s := &GinServer{
		port:   port,
		engine: gin.Default(),
		logger: logger,
	}
	s.registerRoutes(facade)
	return s
}

func (s *GinServer) registerRoutes(facade httphandler.PaymentFacade) {
	s.engine.GET("/test/:name/:name2", func(c *gin.Context) {
		name := c.Param("name")
		name2 := c.Param("name2")
		c.String(http.StatusOK, "test with params: "+name+" and "+name2)
	})

	s.engine.POST("/payments", func(c *gin.Context) {
		s.logger.Info("HTTP Request", "URL", c.Request.URL.String(), "method", c.Request.Method)
		params := make(map[string]string)
		response, statusCode, _ := facade.CreatePayment(httphandler.HttpContext{
			UrlParams: params,
			Body:      "",
		})
		c.String(statusCode, response)
	})

	s.engine.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "post request")
	})

	s.engine.GET("/payments/:id", func(c *gin.Context) {
		s.logger.Info("HTTP Request", "URL", c.Request.URL.String(), "method", c.Request.Method)
		id := c.Param("id")
		response, statusCode, _ := facade.GetPaymentStatus(httphandler.HttpContext{
			UrlParams: map[string]string{"id": id},
			Body:      "",
		})
		c.String(statusCode, response)
	})
}

func (s *GinServer) Serve(ctx context.Context) error {
	return s.engine.Run(":" + s.port)
}
