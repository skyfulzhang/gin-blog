package utils

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func GetLogFile() string {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err = os.Create(fileName); err != nil {
			fmt.Println(err.Error())
			return ""
		}
	}
	return fileName
}

func ZapLogger() *zap.Logger {
	lumberLogger := &lumberjack.Logger{
		Filename:   GetLogFile(),
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     7,
	}
	encoder := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		FunctionKey:    zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	writeSync := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberLogger))
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), writeSync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller(), zap.Development())
	defer logger.Sync()
	// 全局的logger实例，后续在其他包中只需使用zap.L()、zap.S()调用即可
	zap.ReplaceGlobals(logger)
	return logger
}
