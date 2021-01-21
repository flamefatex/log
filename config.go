package log

import "github.com/flamefatex/log/rotation"

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

var currentConfig *Config

// Config 配置项
type Config struct {
	ServiceName    string                   // 服务名称，会带入到日志的service字段中
	Level          string                   // 日志等级，如果不传，则默认是Info
	EnableConsole  bool                     // 是否打印到控制台
	EnableRotation bool                     // 是否记使用日志切割
	RotationConfig *rotation.RotationConfig // 日志切割配置
}

// DefaultConfig 默认配置, 本地开发，联调测试时可以使用此默认配置
func DefaultConfig() *Config {
	return &Config{
		Level:          LevelDebug,
		EnableConsole:  true,
		EnableRotation: false,
	}
}

func CurrentConfig() *Config {
	return currentConfig
}

func CurrentLevel() string {
	return currentConfig.Level
}
