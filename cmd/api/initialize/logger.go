package initialize

import (
	"os"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 初始化hlog日志
func InitHLogger() {
	// 规定好输出日志的目录
	logFilePath := shared.HlogFilePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// 设定好日志的文件名
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	logger := hertzlogrus.NewLogger()
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}
	logger.SetOutput(lumberjackLogger)

	logger.SetLevel(hlog.LevelDebug)
	hlog.SetLogger(logger)
}

// 初始化klog日志
func InitKLogger() {
	// Customizable output directory.
	logFilePath := shared.KlogFilePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	logger := kitexlogrus.NewLogger()
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}
	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(klog.LevelDebug)

	klog.SetLogger(logger)
}
