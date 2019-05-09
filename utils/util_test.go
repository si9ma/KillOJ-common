package utils

import (
	"testing"

	"github.com/si9ma/KillOJ-common/tip"
	"github.com/stretchr/testify/assert"
)

func TestLower1stCharacter(t *testing.T) {
	assert.Equal(t, "helloWorld", Lower1stCharacter("HelloWorld"))
}

func TestCheckEmail(t *testing.T) {
	assert.Equal(t, true, CheckEmail("hellob374@gmail.com"))
}

func TestDeepCopy(t *testing.T) {
	ti := tip.UserAlreadyExistInOrgTip
	tii := tip.Tip{}
	if err := DeepCopy(&tii, &ti); err != nil {
		t.Fatal(err)
	}
	tii["en"] = "test"
	t.Log(ti)
	t.Log(tii)
}
