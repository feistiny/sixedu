package controller

import (
	"github.com/feistiny/sixedu/etc"
	"testing"
)

func Test(t *testing.T) {
	etc.LoadConfig("../etc/config.json")

	t.Run("login", func(t *testing.T) {
		c := NewLogin()
		c.Handle()
	})
	t.Run("show", func(t *testing.T) {
		c := NewShow()
		c.Handle()
	})
}
