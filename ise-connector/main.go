package main

import (
	"ise-connector/internal/logger"
)

func main() {
	log := logger.New(100)

	defer log.Close()

	log.Info("Application starting...")

	log.Info("Application stopped")
}
