package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	logConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
	zap.NewProductionConfig()
}

func Info(msg string, tags ...zap.Field) {
	defer func() {
		err := log.Sync()
		if err != nil {
			//log.Fatal(err.Error()) Error cannot be handled, (library issue)
		}
	}()
	log.Info(msg, tags...)
}

func Error(msg string, err string, tags ...zap.Field) {

	if err != "" {
		tags = append(tags, zap.String("error", err))
	}

	defer func() {
		err := log.Sync()
		if err != nil {
			//log.Fatal(err.Error()) Error cannot be handled, (library issue)
		}
	}()
	log.Error(msg, tags...)
}
