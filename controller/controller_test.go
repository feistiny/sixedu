package controller

import (
	"github.com/feistiny/sixedu/etc"
	"github.com/feistiny/sixedu/util"
	"testing"
)

func Test(t *testing.T) {
	etc.LoadConfig("../etc/config.json")

	t.Run("login", func(t *testing.T) {
		util.SetInput("shineyork", "123456")
		c := NewLogin()
		c.Handle()
	})
	t.Run("show", func(t *testing.T) {
		c := NewShow()
		c.Handle()
	})
}
