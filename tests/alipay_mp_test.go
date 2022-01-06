package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/HaoweiCh/user_agent"
)

func TestAlipayMp(t *testing.T) {
	t.Run("split factor", func(t *testing.T) {
		ua := user_agent.New("Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/18D52 Ariver/1.1.0 AliApp(AP/10.2.51.6000) Nebula WK RVKType(0) AlipayDefined(nt:WIFI,ws:320|504|2.0) AlipayClient/10.2.51.6000 Language/en Region/CN NebulaX/1.0.0")
		name, _ := ua.Browser()
		assert.Equal(t, name, "Alipay MiniProgram")

		t.Log(ua.Prettify())
	})
}
