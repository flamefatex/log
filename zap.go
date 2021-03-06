package log

import (
	"errors"
	"fmt"
	"os"

	"github.com/flamefatex/log/rotation"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	Zap *zap.Logger
	*zap.SugaredLogger
}

func NewZapLogger(c *Config) Logger {
	// 日志等级
	zapLevel := zap.NewAtomicLevel()
	var lvl zapcore.Level
	if err := lvl.Set(c.Level); err != nil {
		err = fmt.Errorf("lvl set err:%w", err)
		panic(err)
	} else {
		zapLevel.SetLevel(lvl)
	}

	// cores里面视情况包含日志切割和console输出的zapcore
	var cores []zapcore.Core

	if c.EnableRotation {
		if c.RotationConfig != nil {
			w := zapcore.AddSync(rotation.NewRotationWriter(c.ServiceName, c.RotationConfig))
			// 无论是生产还是开发环境，都使用相同格式。
			// 方便在开发环境下能获取与生成环境相同格式的日志文件
			fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
			cores = append(cores, zapcore.NewCore(fileEncoder, w, zapLevel))
		}
	}

	if c.EnableConsole {
		// console控制台下只使用开发环境格式
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(consoleEncoder, os.Stdout, zapLevel))
	}

	if len(cores) == 0 {
		z := zap.NewNop()
		return &ZapLogger{Zap: z, SugaredLogger: z.Sugar()}
	}

	core := zapcore.NewTee(cores...)
	// 增加调用行数和stack trace方便定位错误
	z := zap.New(core, zap.AddStacktrace(zap.ErrorLevel), zap.AddCaller(), zap.AddCallerSkip(1), zap.Fields(zap.String("service", c.ServiceName)))
	return &ZapLogger{Zap: z, SugaredLogger: z.Sugar()}
}

func GetZap() (*zap.Logger, error) {
	z, ok := defaultLogger.(*ZapLogger)
	if !ok {
		return nil, errors.New("logger do not init by zap")
	}
	return z.Zap, nil
}
