package log

import (
	"os"
	"testing"

	"github.com/si9ma/KillOJ-common/constants"

	"go.uber.org/zap"
)

func TestLoggerBg(t *testing.T) {
	if err := Init([]string{}, Console); err != nil {
		t.Fatal("Init log fail", err)
		return
	}
	Bg().Info("I'm Info", zap.String("event", "test"), zap.String("level", "info"))
	Bg().Error("I'm Error", zap.String("event", "test"), zap.String("level", "error"))
	Bg().Fatal("I'm Fatal", zap.String("event", "test"), zap.String("level", "fatal"))
}

func TestLoggerBg2File(t *testing.T) {
	if err := Init([]string{"stdout", "/tmp/log.log"}, Console); err != nil {
		t.Fatal("Init log fail", err)
		return
	}
	Bg().Info("I'm Info", zap.String("event", "test"), zap.String("level", "info"))
	Bg().Error("I'm Error", zap.String("event", "test"), zap.String("level", "error"))
	Bg().Fatal("I'm Fatal", zap.String("event", "test"), zap.String("level", "fatal"))
}

func TestLoggerBgDebugModeTrue(t *testing.T) {
	_ = os.Setenv(constants.EnvDebug, "true")
	if err := Init([]string{"/tmp/log.log"}, Console); err != nil {
		t.Fatal("Init log fail", err)
		return
	}

	// should write to file and stdout at the same time
	Bg().Info("I'm Info", zap.String("event", "test"), zap.String("level", "info"))
	Bg().Error("I'm Error", zap.String("event", "test"), zap.String("level", "error"))
	Bg().Fatal("I'm Fatal", zap.String("event", "test"), zap.String("level", "fatal"))
}

func TestLoggerBgDebugModeFalse(t *testing.T) {
	_ = os.Setenv(constants.EnvDebug, "false")
	if err := Init([]string{"/tmp/log.log"}, Console); err != nil {
		t.Fatal("Init log fail", err)
		return
	}

	// should write to file and stdout at the same time
	Bg().Info("I'm Info", zap.String("event", "test"), zap.String("level", "info"))
	Bg().Error("I'm Error", zap.String("event", "test"), zap.String("level", "error"))
	Bg().Fatal("I'm Fatal", zap.String("event", "test"), zap.String("level", "fatal"))
}

func TestLoggerNoManualInit(t *testing.T) {
	Bg().Info("I'm Info", zap.String("event", "test"), zap.String("level", "info"))
	Bg().Error("I'm Error", zap.String("event", "test"), zap.String("level", "error"))
	Bg().Fatal("I'm Fatal", zap.String("event", "test"), zap.String("level", "fatal"))
}

func TestLoggerJson(t *testing.T) {
	if err := Init([]string{}, Json); err != nil {
		t.Fatal("Init log fail", err)
		return
	}

	// should write to file and stdout at the same time
	Bg().Info("I'm Info", zap.String("event", "test"), zap.String("level", "info"))
	Bg().Error("I'm Error", zap.String("event", "test"), zap.String("level", "error"))
	Bg().Fatal("I'm Fatal", zap.String("event", "test"), zap.String("level", "fatal"))
}
