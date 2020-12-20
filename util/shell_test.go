package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("run shell", func(t *testing.T) {
		// command not found
		assert.NotNil(t, RunShell("asdf"))
		assert.Nil(t, RunShell("wc -l ../data/user.sql.orig | grep '^2 '"))
	})
}
