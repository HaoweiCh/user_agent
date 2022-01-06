package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/HaoweiCh/user_agent"
)

func TestWechatMp(t *testing.T) {
	t.Run("split factor", func(t *testing.T) {
		ua := user_agent.New("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat")
		name, _ := ua.Browser()
		assert.Equal(t, name, "Wechat MiniProgram")

		t.Log(ua.Prettify())
	})
}
