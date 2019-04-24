package asyncjob

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/brokers/iface"

	machlog "github.com/RichardKnop/machinery/v1/log"
	"github.com/si9ma/KillOJ-common/log"
)

var machineryServer *machinery.Server

// manual init
func Init(cfg Config) (err error) {
	machineryCfg := cfg.toMachineryCfg()

	var machineryBroker iface.Broker
	if machineryBroker, err = machinery.BrokerFactory(machineryCfg); err != nil {
		return err
	}
	machineryBackend := NewBlankBackend()

	machineryServer = machinery.NewServerWithBrokerBackend(machineryCfg, machineryBroker, machineryBackend)
	machlog.Set(ZLogger{}) // custom logger
	return
}

// return the server
func Server() *machinery.Server {
	if machineryServer == nil {
		log.Bg().Error("Please Init Machinery Server Before use it")
	}

	return machineryServer
}
