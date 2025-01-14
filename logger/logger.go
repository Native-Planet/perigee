package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	once sync.Once
	log  *zap.Logger
)

func Init() error {
	var err error
	once.Do(func() {
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		config.OutputPaths = []string{"stdout"}

		log, err = config.Build()
		if err != nil {
			return
		}
		zap.ReplaceGlobals(log)
	})
	return err
}

func GetLogger() *zap.Logger {
	return zap.L()
}

func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}
