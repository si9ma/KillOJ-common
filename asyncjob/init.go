package asyncjob

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/si9ma/KillOJ-common/log"
	"go.uber.org/zap"
)

var machineryServer *machinery.Server

// auto init
func init() {
	log.Bg().Warn("auto init machinery server,amqp broker may wrong, Please init machinery manual",
		zap.String("config", defaultConfig().String()))
	machineryServer, _ = machinery.NewServer(defaultConfig().toMachineryCfg())
}

// manual init
func Init(cfg Config) (err error) {
	machineryServer, err = machinery.NewServer(cfg.toMachineryCfg())
	return
}

// return the server
func Server() *machinery.Server {
	return machineryServer
}
