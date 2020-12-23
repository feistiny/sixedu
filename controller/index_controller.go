package controller

import (
	"github.com/astaxie/beego/logs"
	"github.com/feistiny/sixedu/router"
)

type indexController struct {
}

func NewIndex() Controller {
	return &indexController{}
}

func (lc *indexController) Handle() (success bool, routes router.NextRoutes) {
	logs.Debug("index start")
	success = true
	routes = router.NextRoutes{router.LoginRoute, router.RegRoute}
	return
}
