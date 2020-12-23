package controller

import (
	"github.com/feistiny/sixedu/model"
	"github.com/feistiny/sixedu/router"
	"github.com/feistiny/sixedu/util"
	"strconv"
)

type registerController struct {
}

// NewReg
func NewReg() Controller {
	return &registerController{}
}

func (rc *registerController) Handle() (success bool, routes router.NextRoutes) {
	println("输入你需要注册的用户信息 username,password,age,sex")
	u := model.NewUser()
	username := util.GetInput("输入你的用户名")
	u.SetUsername(username)
	pwd := util.GetInput("输入你的密码")
	pwd2 := util.GetInput("确认密码")
	if pwd != pwd2 {
		println("两次输入的密码不一致")
		success = false
		return
	}
	u.SetPassword(pwd)
	ageRaw := util.GetInput("输入你的年龄")
	age, err := strconv.ParseInt(ageRaw, 0, 0)
	if err != nil {
		println("年龄输入错误")
		success = false
		return
	}
	u.SetAge(int(age))
	sex := util.GetInput("输入你的性别")
	u.SetSex(sex)
	if !u.Save() {
		success = false
		return
	}

	success = true
	routes = router.NextRoutes{router.LoginRoute, router.RegRoute}
	return
}
