package main

import (
	"context"
	"log"
	"payway/internal/app"
	"payway/internal/controller/http"
	"payway/internal/webserver"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	paymentRepo := app.NewPaymentService()

	facade := http.NewPaymentFacade(paymentRepo)

	// handlerType := os.Getenv("HTTP_HANDLER")
	// port := os.Getenv("PORT")

	handlerType := "gin"
	// handlerType := "chi"
	port := "8080"

	server := webserver.NewWebServer(ctx, handlerType, port)

	routes := http.GetRoutes(facade)

	server.RegisterRoutes(routes)

	log.Printf("Starting server on :%s", port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// curl -X POST http://localhost:80/payments \
//   -H "Content-Type: application/json" \
//   -d '{"user":234,"amount":"445.67","currency":"RUB","idempotencyKey":"abcde"}'
