package log

import (
	"errors"
	"os"

	"github.com/flamefatex/log/impl"
	"github.com/flamefatex/log/impl/w_zap"
	"go.uber.org/zap"
)

var defaultLogger impl.Logger

type NewLoggerFunc func(*Config) impl.Logger

func init() {
	// 环境变量支持
	serviceName := os.Getenv("SERVICE_NAME")
	SetServiceName(serviceName)
	mode := os.Getenv("ENV_MODE")
	SetMode(mode)

	defaultLogger = w_zap.NewZapLogger(DefaultConfig())
}

func NewLogger(f NewLoggerFunc, c *Config) impl.Logger {
	return f(c)
}

func InitLogger(f NewLoggerFunc, c *Config) {
	defaultLogger = f(c)
}

func SetLogger(logger impl.Logger) {
	defaultLogger = logger
}

func L() impl.Logger {
	return defaultLogger
}

func GetZap() (*zap.Logger, error) {
	z, ok := defaultLogger.(*w_zap.ZapLogger)
	if !ok {
		return nil, errors.New("logger do not init by zap")
	}
	return z.Zap, nil
}

// Debug package-zapLevel convenience method.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf package-zapLevel convenience method.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Info package-zapLevel convenience method.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof package-zapLevel convenience method.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warn package-zapLevel convenience method.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf package-zapLevel convenience method.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Error package-zapLevel convenience method.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf package-zapLevel convenience method.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatal package-zapLevel convenience method.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf package-zapLevel convenience method.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Panic package-zapLevel convenience method.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Panicf package-zapLevel convenience method.
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}
