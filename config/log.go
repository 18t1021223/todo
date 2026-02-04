package config

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(env string) {
	logDir := "./logs"
	logErrorPath := fmt.Sprintf("%s/error.log", logDir)
	logPath := fmt.Sprintf("%s/app.log", logDir)

	// Create log folder
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}

	var config zap.Config

	if env == "prod" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "ts"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		// Development: Console format, Debug level
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Output files
	config.OutputPaths = []string{"stdout", logPath}
	config.ErrorOutputPaths = []string{"stderr", logErrorPath}

	var err error
	Logger, err = config.Build(
		zap.AddCaller(),                       // Add file info to into line number
		zap.AddStacktrace(zapcore.ErrorLevel), // Stack trace for error or more
	)
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(Logger)
}
