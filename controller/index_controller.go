package controller

import (
	"github.com/astaxie/beego/logs"
)

type indexController struct {
}

func NewIndex() Controller {
	return &indexController{}
}

func (lc *indexController) Handle() (success bool, routes NextRoutes) {
	logs.Debug("index start")
	success = true
	routes = NextRoutes{LoginRoute, RegRoute}
	return
}
