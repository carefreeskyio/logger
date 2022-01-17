package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"time"
)

type Option struct {
	Path         string
	RotationTime int
	Level        string
	WithMaxAge   int
}

func InitLogger(o *Option) {
	writer, err := rotatelogs.New(
		o.Path+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(o.Path),
		rotatelogs.WithRotationTime(time.Duration(o.RotationTime)*time.Hour),
		rotatelogs.WithMaxAge(time.Duration(o.WithMaxAge*24)*time.Hour),
	)
	fmt.Println(err)
	log.SetOutput(writer)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	if level, err := log.ParseLevel(o.Level); err != nil {
		log.Fatalf("init log failed: err=%v", err)
		log.SetLevel(level)
	}
	log.SetReportCaller(true)
}
