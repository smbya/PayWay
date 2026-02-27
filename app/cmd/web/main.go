package main

import (
	"context"
	"payway/internal/app"
	"payway/internal/config"
	"payway/internal/controller/http"
	"payway/internal/controller/webserver"
	"payway/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	cfg := config.CreateConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.Db.GetConnectString())
	if err != nil {
		cfg.Logger.Error("Failed to connect to database", "Error", err)
		return
	}
	defer pool.Close()

	cfg.Logger.Info("Connected to database")

	repository := repository.NewRepository(pool)

	paymentService := app.NewPaymentService(repository, cfg.Logger)

	facade := http.NewPaymentFacade(paymentService, cfg.Logger)

	server := webserver.NewWebServer(cfg.App.HandlerType, cfg.App.Port, cfg.Logger, facade)

	cfg.Logger.Info("Starting server", "port", cfg.App.Port)

	if err := server.Serve(ctx); err != nil {
		cfg.Logger.Error("Server failed to start", "error", err)
		return
	}
}

// curl -X POST http://localhost:80/payments \
//   -H "Content-Type: application/json" \
//   -d '{"user":234,"amount":"445.67","currency":"RUB","idempotencyKey":"abcde"}'
