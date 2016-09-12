package log

import (
	"fmt"

	"github.com/k4s/tea/log"
	"github.com/robfig/cron"
)

func init() {
	c := cron.New()
	c.AddFunc("0 0 0 * * ?", func() {
		fmt.Println("new log")
		var err error
		Logger, err = log.New("debug", "./log/logdata/")
		if err != nil {
			log.Debug("%v", err)
		}
	})
	c.Start()
}

var Logger, _ = log.New("debug", "./log/logdata/")

var gLogger, _ = log.New("debug", "")

func Debug(format string, a ...interface{}) {
	gLogger.Debug(format, a...)
}

func Release(format string, a ...interface{}) {
	gLogger.Release(format, a...)
}

func Error(format string, a ...interface{}) {
	gLogger.Error(format, a...)
}

func Fatal(format string, a ...interface{}) {
	gLogger.Fatal(format, a...)
}

func Close() {
	gLogger.Close()
}
