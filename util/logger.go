package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var (
	LogRus *logrus.Logger
)

func InitLog(configFile string) {
	//viper := CreateConfig(configFile)
	LogRus = logrus.New()
	//switch strings.ToLower(viper.("level")) {
	//case "debug":
	//	LogRus.SetLevel(logrus.DebugLevel)
	//}
	switch configFile {
	case "debug":
		LogRus.SetLevel(logrus.DebugLevel)
	case "info":
		LogRus.SetLevel(logrus.InfoLevel)
	case "warn":
		LogRus.SetLevel(logrus.InfoLevel)
	case "error":
		LogRus.SetLevel(logrus.ErrorLevel)
	case "panic":
		LogRus.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("invalid log level %s", "DeBug"))
	}
	LogRus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//logFile := "222"
	//fout, err := rotatelogs.New(
	//
	//	)
	//if err != nil {
	//	panic(err)
	//}
	//LogRus.SetOutput(fout)
	LogRus.SetReportCaller(true)
}
