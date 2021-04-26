// nolint: gochecknoinits, gochecknoglobals
package log

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerKeyType int

const loggerKey loggerKeyType = iota

var (
	atom   zap.AtomicLevel
	logger *zap.SugaredLogger
)

func init() {
	atom = zap.NewAtomicLevel()

	l := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atom,
	))

	logger = l.Sugar()
}

// SetLevel dinamically changes the log level.
func SetLevel(lvl string) {
	atom.UnmarshalText([]byte(lvl))
}

// NewContext creates a new logger with fields.
func NewContext(ctx context.Context, fields ...interface{}) context.Context {
	return context.WithValue(ctx, loggerKey, WithContext(ctx).With(fields...))
}

// NewContextLog creates a new logger context from another logger instance.
func NewContextLog(ctx context.Context, l *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

// WithContext gets the existing logger from context. If not present returns the default.
func WithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return logger
	}

	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.SugaredLogger); ok {
		return ctxLogger
	}

	return logger
}
