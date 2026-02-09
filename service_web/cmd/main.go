package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Service web start")

	// Создаем канал для обработки сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ждем сигнала для завершения
	<-sigChan
	log.Println("Received shutdown signal")
}
