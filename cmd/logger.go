package cmd

import (
	"github.com/alextuan1024/requiem/log"
	"go.uber.org/zap"
)

func logger() *zap.SugaredLogger {
	return log.Logger()
}
