package util

import (
	"github.com/astaxie/beego/config"
	"github.com/feistiny/sixedu/etc"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// 跟真实环境相关的测试, 默认跳过, 如果指定了环境变量REAL, 则运行
func SkipReal(t *testing.T) {
	if os.Getenv("REAL") == "" {
		t.Skip("跳过真实环境的测试")
	}
}

func FakeConfig(t *testing.T) {
	etc.Cfg = config.NewFakeConfig()
	assert.Nil(t, etc.Cfg.Set("SQL_SUFFIX", ".sql.test"))
	assert.Nil(t, etc.Cfg.Set("SQL_PATH", "../data/"))
}
