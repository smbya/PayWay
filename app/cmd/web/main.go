package main

import (
	"context"
	"log"
	"payway/internal/app"
	"payway/internal/config"
	"payway/internal/controller/http"
	"payway/internal/repository"
	"payway/internal/webserver"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	cfg := config.CreateConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.Db.GetConnectString())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repository := repository.NewRepositury(pool)

	paymentService := app.NewPaymentService(repository)

	facade := http.NewPaymentFacade(paymentService)

	server := webserver.NewWebServer(ctx, cfg.App.HandlerType, cfg.App.Port)

	routes := http.GetRoutes(facade)

	server.RegisterRoutes(routes)

	log.Printf("Starting server on :%s", cfg.App.Port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// curl -X POST http://localhost:80/payments \
//   -H "Content-Type: application/json" \
//   -d '{"user":234,"amount":"445.67","currency":"RUB","idempotencyKey":"abcde"}'
