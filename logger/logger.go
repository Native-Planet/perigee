package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	once   sync.Once
	logger *zap.Logger
)

func Init() error {
	var err error
	once.Do(func() {
		logger, err = zap.NewProduction()
		if err != nil {
			return
		}
		zap.ReplaceGlobals(logger)
	})
	return err
}

func GetLogger() *zap.Logger {
	return logger
}

func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}
