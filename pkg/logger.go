package logger

import (
	"log"

	"go.uber.org/zap"
)

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}
	zap.ReplaceGlobals(logger)
}
