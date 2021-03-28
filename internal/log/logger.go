package log

import (
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var Logger *logrus.Entry

func init() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(cwd, "logs", "watchdog.log")

	// 下面配置日志每隔 30 分钟轮转一个新文件，保留最近 12 小时的日志文件，多余的自动清理掉。
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Hour*12),
		rotatelogs.WithRotationTime(time.Minute*30),
	)
	log.SetOutput(writer)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	Logger = log.WithFields(log.Fields{})
}
