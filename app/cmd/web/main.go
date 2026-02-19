package main

import (
	"context"
	"log"
	"payway/internal/app"
	"payway/internal/controller/http"
	"payway/internal/repository"
	"payway/internal/webserver"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.New(ctx, "postgres://user:password@db:5432/payway?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repository := repository.NewRepositury(pool)

	paymentService := app.NewPaymentService(repository)

	facade := http.NewPaymentFacade(paymentService)

	// handlerType := os.Getenv("HTTP_HANDLER")
	// port := os.Getenv("PORT")

	handlerType := "gin"
	// handlerType := "chi"
	port := "80"

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
