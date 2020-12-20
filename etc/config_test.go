package etc

import (
	"github.com/astaxie/beego/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigFileNotExist(t *testing.T) {
	err := LoadConfig("non-existent.file")
	assert.NotNil(t, err)
}

func TestGetFromRealFile(t *testing.T) {
	err := LoadConfig("./config.json.example")
	assert.Nil(t, err)
	v := Cfg.String("key")
	assert.Equal(t, "value", v)
}

func TestType(t *testing.T) {
	Cfg = config.NewFakeConfig()
	var err error
	err = Cfg.Set("string", "string")
	assert.Nil(t, err)
	err = Cfg.Set("bool", "true")
	assert.Nil(t, err)
	t.Run("string", func(t *testing.T) {
		v := Cfg.String("string")
		assert.Equal(t, "string", v)
	})
	t.Run("bool", func(t *testing.T) {
		v, err := Cfg.Bool("bool")
		assert.Nil(t, err)
		assert.Equal(t, true, v)
	})
}
