package utils

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

const (
	context = "[Gateway]"
)

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction()

	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	defer Logger.Sync()
}

func LoggerInfo(msg string) {
	Logger.Info(fmt.Sprintf("%v %v", context, msg))
}

func LoggerInfof(msg string, args ...interface{}) {
	Logger.Sugar().Infof(fmt.Sprintf("%v %v", context, msg), args)
}

func LoggerError(msg string) {
	Logger.Error(fmt.Sprintf("%v %v", context, msg))
}

func LoggerErrorf(msg string, args ...interface{}) {
	Logger.Sugar().Errorf(fmt.Sprintf("%v %v", context, msg), args)
}

func LoggerFatalf(msg string, args ...interface{}) {
	Logger.Sugar().Fatalf(fmt.Sprintf("%v %v", context, msg), args)
}

func LoggerFatal(msg string) {
	Logger.Fatal(fmt.Sprintf("%v %v", context, msg))
}
