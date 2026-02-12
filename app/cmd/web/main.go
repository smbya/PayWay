package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payment struct {
	UserId         int    `json:"user"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	IdempotencyKey string `json:"idempotencyKey"`
}

func main() {
	log.Println("Service web start")

	// Создаем канал для обработки сигналов
	// 	sigChan := make(chan os.Signal, 1)
	// 	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 	// Ждем сигнала для завершения
	// 	<-sigChan
	// 	log.Println("Received shutdown signal")

	http.HandleFunc("POST /payments", createPayment)
	http.HandleFunc("GET /payment", getPayment)

	log.Println("Server started on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func getPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("hello"))
}

// curl -X POST http://localhost:80/payments \
//   -H "Content-Type: application/json" \
//   -d '{"user":234,"amount":"445.67","currency":"RUB","idempotencyKey":"abcde"}'
