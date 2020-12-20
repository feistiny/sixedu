package model

import (
	"fmt"
	"github.com/feistiny/sixedu/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	// setup
	util.FakeConfig(t)

	assert.Nil(t, util.RunShell("wc -l ../data/user.sql.orig | grep '^2 '"))

	t.Run("测试 model 写入 sql", func(t *testing.T) {
		assert.Nil(t, util.RunShell("cp ../data/user.sql.orig ../data/user.sql.test"))
		u := &User{
			username: "test",
			password: "test",
			age:      3,
			sex:      "男",
		}
		assert.Nil(t, rwdata("user", u))
		assert.Nil(t, util.RunShell("wc -l ../data/user.sql.test | grep '^3 '"))
		assert.Nil(t, util.RunShell("grep %s ../data/user.sql.test", u.username))
	})
	t.Run("测试 model 读出 sql", func(t *testing.T) {
		assert.Nil(t, util.RunShell("cp ../data/user.sql.orig ../data/user.sql.test"))
		datas, err := rfdata("user", "shineyork")
		assert.Nil(t, err)
		for _, m := range datas {
			fmt.Println(m.ToStringSlice())
			// fmt.Println(m)
		}
	})
}
