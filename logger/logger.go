package logger

import (
	"go.uber.org/zap"
	"log"
)

var Logger *zap.Logger

func init() {
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Unable to initialize logger: %v", err)
	}
}
