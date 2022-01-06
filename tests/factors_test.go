package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/HaoweiCh/user_agent"
)

func TestFactors(t *testing.T) {
	t.Run("dev", func(t *testing.T) {
		factors := user_agent.Factors{}
		assert.NoError(t, factors.Load("../factors.json"))
		assert.NoError(t, factors.Generate("../factors_raw.txt"))
		assert.NoError(t, factors.Dump("../factors.json"))
	})
}
