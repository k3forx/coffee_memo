package logger

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/config"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func Init() (func(), error) {
	var err error
	if config.IsProduction() {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}
	fn := func() {
		_ = logger.Sync()
	}
	return fn, err
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(ctx context.Context, err error, fields ...zap.Field) {
	logger.Error(err.Error(), fields...)
}
