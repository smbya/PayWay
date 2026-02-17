package gin

import (
	"context"
	"log"
	"payway/internal/controller/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	ctx    context.Context
	port   string
	engine *gin.Engine
}

func replacePathVariables(s string) string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	return re.ReplaceAllString(s, `:$1`)
}

func NewGinServer(ctx context.Context, port string) *GinServer {
	return &GinServer{
		ctx:    ctx,
		port:   port,
		engine: gin.Default(),
	}
}

func (s *GinServer) RegisterRoutes(routes []http.Route) {
	for _, route := range routes {

		ginPathWithVariables := replacePathVariables(route.Path)

		log.Print(route.Method, ginPathWithVariables)

		s.engine.Handle(route.Method, ginPathWithVariables, func(c *gin.Context) {

			params := make(map[string]string)
			for _, p := range c.Params {
				params[p.Key] = p.Value
			}

			httpContext := http.HttpContext{
				UrlParams: params,
				Body:      "", //r.Body,
			}

			response, statusCode, _ := route.Handler(httpContext)
			c.String(statusCode, response)
		})
	}
}

func (s *GinServer) Run() error {
	return s.engine.Run(":" + s.port)
}
