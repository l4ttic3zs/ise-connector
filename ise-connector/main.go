package main

import (
	"time"

	"ise-connector/internal/logger"
)

func main() {
	log := logger.New(100)

	defer log.Close()

	log.Info("Application starting...")

	time.Sleep(5 * time.Second)

	log.Info("Application stopped")
}
