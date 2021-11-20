package main

import (
	"github.com/danilocordeiro/banking/app"
	"github.com/danilocordeiro/banking/logger"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}