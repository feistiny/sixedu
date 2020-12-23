package controller

import (
	"github.com/astaxie/beego/logs"
	"github.com/feistiny/sixedu/model"
	"github.com/feistiny/sixedu/util"
)

type loginController struct {
}

func NewLogin() controller {
	return &loginController{}
}

func (lc *loginController) Handle() (success bool, routes nextRoutes) {
	logs.Debug("login start")
	u := model.NewUser()
	datas, err := u.All()
	if err != nil {
		logs.Error("登录的账号列表获取失败")
		success = false
		return
	}

	if len(datas) <= 0 {
		success = false
		return
	}

	username := util.GetInput("输入你的用户名")
	pwd := util.GetInput("输入你的密码")

	var ok bool
	if _, ok = datas[username]; !ok {
		println("用户不存在", username)
		success = false
		return
	}

	me := datas[username].(*model.User)
	println("username", me.GetUsername())
	println("password", me.GetPassword())

	if pwd != me.GetPassword() {
		success = false
		return
	}

	success = true
	return
}
