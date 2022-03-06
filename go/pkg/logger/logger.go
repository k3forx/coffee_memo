package logger

import (
	"context"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func Init() (func(), error) {
	var err error
	logger, err := zap.NewProduction()
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
