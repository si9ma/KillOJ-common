package tip

import (
	"os"
	"testing"

	"github.com/si9ma/KillOJ-common/constants"
)

func TestTip_String(t *testing.T) {
	if err := os.Setenv(constants.EnvLang, "zh"); err != nil {
		t.Fatal(err)
	}

	t.Log(RunSuccessTip.String())
}
