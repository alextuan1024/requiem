package log

import "go.uber.org/zap"

var slogger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	slogger = logger.Sugar()
}

func Logger() *zap.SugaredLogger {
	return slogger
}
