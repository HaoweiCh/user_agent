package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/HaoweiCh/user_agent"
)

func TestFactors(t *testing.T) {
	t.Run("check", func(t *testing.T) {
		t.Logf("%s", user_agent.DefaultFactors)
	})
	t.Run("generate", func(t *testing.T) {
		factors := user_agent.Factors{}
		assert.NoError(t, factors.Load("../factors.json"))
		assert.NoError(t, factors.Generate("../factors_raw.txt"))
		assert.NoError(t, factors.Dump("../factors.json"))
	})
}
