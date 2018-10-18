package log

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

type loggerKeyType int

const loggerKey loggerKeyType = iota

var logger *log.Logger

func init() {
	log.SetHandler(cli.New(os.Stdout))
	if l, ok := log.Log.(*log.Logger); ok {
		logger = l
	}
}

// NewContext returns a context that has a logrus logger
func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerKey, WithContext(ctx))
}

// WithContext returns a logrus logger from the context
func WithContext(ctx context.Context) *log.Logger {
	if ctx == nil {
		return logger
	}

	if ctxLogger, ok := ctx.Value(loggerKey).(*log.Logger); ok {
		return ctxLogger
	}

	return logger
}
