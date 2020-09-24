package log

import "github.com/flamefatex/log/rotation"

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"

	ModeDebug      = "debug"
	ModeProduction = "production"
)

var defaultConfig = &Config{
	Mode:           ModeDebug,
	Level:          LevelDebug,
	EnableConsole:  true,
	EnableRotation: false,
}

// Config 配置项
type Config struct {
	ServiceName    string                   // 服务名称，会带入到日志的service字段中
	Mode           string                   // 模式
	Level          string                   // 日志等级，如果不传，则默认是Info
	EnableConsole  bool                     // 是否打印到控制台
	EnableRotation bool                     // 是否记使用日志切割
	RotationConfig *rotation.RotationConfig // 日志切割配置
}

// DefaultConfig 默认配置, 本地开发，联调测试时可以使用此默认配置
func DefaultConfig() *Config {
	return defaultConfig
}

func SetServiceName(name string) {
	defaultConfig.ServiceName = name
}

func SetMode(value string) {
	switch value {
	case ModeDebug, "":
		defaultConfig.Mode = ModeDebug
		defaultConfig.Level = LevelDebug
	case ModeProduction:
		defaultConfig.Mode = ModeProduction
		defaultConfig.Level = LevelInfo
	default:
		panic("env mode unknown: " + value)
	}
}

// Mode returns currently env mode.
func Mode() string {
	return defaultConfig.Mode
}
