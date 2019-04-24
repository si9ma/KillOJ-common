package asyncjob

import (
	"fmt"

	"github.com/si9ma/KillOJ-common/log"
)

type ZLogger struct{}

func (l ZLogger) Print(args ...interface{}) {
	log.Bg().Info(fmt.Sprint(args...))
}

func (l ZLogger) Printf(s string, args ...interface{}) {
	log.Bg().Info(fmt.Sprintf(s, args))
}

func (l ZLogger) Println(args ...interface{}) {
	log.Bg().Info(fmt.Sprintln(args))
}

func (l ZLogger) Fatal(args ...interface{}) {
	log.Bg().Error(fmt.Sprint(args...))
}

func (l ZLogger) Fatalf(s string, args ...interface{}) {
	log.Bg().Error(fmt.Sprintf(s, args))
}

func (l ZLogger) Fatalln(args ...interface{}) {
	log.Bg().Error(fmt.Sprintln(args))
}

func (l ZLogger) Panic(args ...interface{}) {
	log.Bg().Fatal(fmt.Sprint(args...))
}

func (l ZLogger) Panicf(s string, args ...interface{}) {
	log.Bg().Fatal(fmt.Sprintf(s, args))
}

func (l ZLogger) Panicln(args ...interface{}) {
	log.Bg().Fatal(fmt.Sprintln(args))
}
