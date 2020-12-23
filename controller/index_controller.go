package controller

import (
	"github.com/astaxie/beego/logs"
)

type indexController struct {
}

func NewIndex() controller {
	return &indexController{}
}

func (lc *indexController) Handle() (success bool, routes nextRoutes) {
	logs.Debug("index start")
	success = true
	routes = nextRoutes{LoginRoute, RegRoute}
	return
}
