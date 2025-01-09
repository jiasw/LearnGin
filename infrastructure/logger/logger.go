package logger

import (
	"github.com/sirupsen/logrus"

	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetOutput(&lumberjack.Logger{
		Filename:   "logs/application.log",
		MaxSize:    10,    // 每个日志文件的最大尺寸（MB）
		MaxBackups: 3,     // 保留旧的日志文件个数
		MaxAge:     28,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	})
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

}

func Info(msg string) {
	Logger.Info(msg)
}

func Error(msg string) {
	Logger.Error(msg)
}

func Debug(msg string) {
	Logger.Debug(msg)
}

func Warning(msg string) {
	Logger.Warning(msg)
}

func Fatal(msg string) {
	Logger.Fatal(msg)
}
