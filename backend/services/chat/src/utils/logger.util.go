package utils

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

const (
	loggerContext = "[Chat]"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction()

	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	defer Logger.Sync()
}

func LoggerInfo(msg string) {
	Logger.Info(fmt.Sprintf("%v %v", loggerContext, msg))
}

func LoggerInfof(msg string, args ...interface{}) {
	Logger.Sugar().Infof(fmt.Sprintf("%v %v", loggerContext, msg), args)
}

func LoggerFatalf(msg string, args ...interface{}) {
	Logger.Sugar().Fatalf(fmt.Sprintf("%v %v", loggerContext, msg), args)
}

func LoggerFatal(msg string) {
	Logger.Fatal(fmt.Sprintf("%v %v", loggerContext, msg))
}

func LoggerErrorf(msg string, args ...interface{}) {
	Logger.Sugar().Errorf(fmt.Sprintf("%v %v", loggerContext, msg), args)
}

func LoggerError(msg string) {
	Logger.Error(fmt.Sprintf("%v %v", loggerContext, msg))
}
