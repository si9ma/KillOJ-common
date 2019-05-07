package lang

import (
	"os"
	"testing"

	"github.com/si9ma/KillOJ-common/constants"
)

func TestGetLangFromEnv(t *testing.T) {
	if err := os.Setenv(constants.EnvLang, "zh"); err != nil {
		t.Fatal(err)
	}

	t.Log(GetLangFromEnv())
}
