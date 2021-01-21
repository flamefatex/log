package log

var defaultLogger Logger

type NewLoggerFunc func(*Config) Logger

func init() {
	InitLogger(NewZapLogger, DefaultConfig())
}

func NewLogger(f NewLoggerFunc, c *Config) Logger {
	return f(c)
}

func InitLogger(f NewLoggerFunc, c *Config) {
	currentConfig = c
	defaultLogger = f(c)
}

// Debug package-zapLevel convenience method.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf package-zapLevel convenience method.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Debugw package-zapLevel convenience method.
func Debugw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Debugw(msg, keysAndValues...)
}

// Info package-zapLevel convenience method.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof package-zapLevel convenience method.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Infow package-zapLevel convenience method.
func Infow(msg string, keysAndValues ...interface{}) {
	defaultLogger.Infow(msg, keysAndValues...)
}

// Warn package-zapLevel convenience method.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf package-zapLevel convenience method.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Warnw package-zapLevel convenience method.
func Warnw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Warnw(msg, keysAndValues...)
}

// Error package-zapLevel convenience method.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Error package-zapLevel convenience method.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Errorw package-zapLevel convenience method.
func Errorw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Errorw(msg, keysAndValues...)
}

// Fatal package-zapLevel convenience method.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf package-zapLevel convenience method.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Fatalw package-zapLevel convenience method.
func Fatalw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Fatalw(msg, keysAndValues...)
}

// Panic package-zapLevel convenience method.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Panicf package-zapLevel convenience method.
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

// Panicw package-zapLevel convenience method.
func Panicw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Panicw(msg, keysAndValues...)
}
