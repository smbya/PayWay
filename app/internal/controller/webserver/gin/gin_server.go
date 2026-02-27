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

func NewGinServer(port string, logger *slog.Logger, handler *httphandler.PaymentHandler) *GinServer {
	s := &GinServer{
		port:   port,
		engine: gin.Default(),
		logger: logger,
	}
	s.registerRoutes(handler)
	return s
}

func (s *GinServer) registerRoutes(handler *httphandler.PaymentHandler) {
	s.engine.POST("/payments", func(c *gin.Context) {
		s.logger.Info("HTTP Request", "URL", c.Request.URL.String(), "method", c.Request.Method)

		var req httphandler.CreatePaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payment, err := handler.CreatePayment(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, payment)
	})

	s.engine.GET("/payments/:id", func(c *gin.Context) {
		s.logger.Info("HTTP Request", "URL", c.Request.URL.String(), "method", c.Request.Method)
		id := c.Param("id")

		payment, err := handler.GetPaymentByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, payment)
	})
}

func (s *GinServer) Serve(ctx context.Context) error {
	return s.engine.Run(":" + s.port)
}
