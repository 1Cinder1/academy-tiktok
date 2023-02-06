package boot

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_tiktok/app/global"
	"os"
	"time"
)

func LoggerSetup() {
	dynamicLevel := zap.NewAtomicLevel()

	switch global.Config.Logger.LogLevel {
	case "debug":
		dynamicLevel.SetLevel(zap.DebugLevel)
	case "info":
		dynamicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		dynamicLevel.SetLevel(zap.WarnLevel)
	case "error":
		dynamicLevel.SetLevel(zap.ErrorLevel)
	}

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	})
	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, os.Stdout, dynamicLevel),
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   global.Config.Logger.SavePath,
				MaxSize:    global.Config.Logger.MaxSize,
				MaxAge:     global.Config.Logger.MaxAge,
				MaxBackups: global.Config.Logger.MaxBackups,
				LocalTime:  true,
				Compress:   global.Config.Logger.IsCompres,
			}),
			dynamicLevel,
		),
	}
	global.Logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	defer func(Logger *zap.Logger) {
		_ = Logger.Sync()
	}(global.Logger)

	global.Logger.Info("initialize logger successfully!")
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
}
func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,                            // 日志保存位置
		MaxSize:    global.Config.Logger.MaxSize,    // 日志文件最大大小 (MB)
		MaxBackups: global.Config.Logger.MaxBackups, // 日志文件备份数量
		MaxAge:     global.Config.Logger.MaxAge,     // 日志保存时间
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
