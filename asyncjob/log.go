package asyncjob

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/si9ma/KillOJ-common/utils"

	"github.com/si9ma/KillOJ-common/log"
)

var proxyLog zap.Field = zap.String("proxy", "machinery")

type ZLogger struct{}

func (l ZLogger) Print(args ...interface{}) {
	log.Bg().Info(fmt.Sprint(args...), proxyLog)
}

func (l ZLogger) Printf(s string, args ...interface{}) {
	log.Bg().Info(fmt.Sprintf(s, args...), proxyLog)
}

func (l ZLogger) Println(args ...interface{}) {
	log.Bg().Info(fmt.Sprintln(args...), proxyLog)
}

func (l ZLogger) Fatal(args ...interface{}) {
	log.Bg().Error(fmt.Sprint(args...), proxyLog)
}

func (l ZLogger) Fatalf(s string, args ...interface{}) {
	log.Bg().Error(fmt.Sprintf(s, args...), proxyLog)
}

func (l ZLogger) Fatalln(args ...interface{}) {
	log.Bg().Error(fmt.Sprintln(args...), proxyLog)
}

func (l ZLogger) Panic(args ...interface{}) {
	log.Bg().Panic(fmt.Sprint(args...), proxyLog)
}

func (l ZLogger) Panicf(s string, args ...interface{}) {
	log.Bg().Panic(fmt.Sprintf(s, args...), proxyLog)
}

func (l ZLogger) Panicln(args ...interface{}) {
	log.Bg().Panic(fmt.Sprintln(args...), proxyLog)
}

type DebugLogger struct{}

func (d DebugLogger) Print(args ...interface{}) {
	if utils.IsDebug() {
		log.Bg().Info(fmt.Sprint(args...), proxyLog)
	}
}

func (d DebugLogger) Printf(s string, args ...interface{}) {
	if utils.IsDebug() {
		log.Bg().Info(fmt.Sprintf(s, args...), proxyLog)
	}
}

func (d DebugLogger) Println(args ...interface{}) {
	if utils.IsDebug() {
		log.Bg().Info(fmt.Sprintln(args...), proxyLog)
	}
}

func (d DebugLogger) Fatal(args ...interface{}) {
	if utils.IsDebug() {
		log.Bg().Error(fmt.Sprint(args...), proxyLog)
	}
}

func (d DebugLogger) Fatalf(s string, args ...interface{}) {
	if utils.IsDebug() {
		log.Bg().Error(fmt.Sprintf(s, args...), proxyLog)
	}
}

func (d DebugLogger) Fatalln(args ...interface{}) {
	if utils.IsDebug() {
		log.Bg().Error(fmt.Sprintln(args...), proxyLog)
	}
}

func (d DebugLogger) Panic(args ...interface{}) {
	if utils.IsDebug() {
		log.Bg().Panic(fmt.Sprint(args...), proxyLog)
	}
}

func (d DebugLogger) Panicf(s string, args ...interface{}) {
	if utils.IsDebug() {
		log.Bg().Panic(fmt.Sprintf(s, args...), proxyLog)
	}
}

func (d DebugLogger) Panicln(args ...interface{}) {
	if utils.IsDebug() {
		log.Bg().Panic(fmt.Sprintln(args...), proxyLog)
	}
}
