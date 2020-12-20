package controller

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/feistiny/sixedu/model"
	"github.com/feistiny/sixedu/util"
	"reflect"
	"strings"
)

type showController struct {
}

func NewShow() *showController {
	return &showController{}
}

func (lc *showController) Handle() bool {
	logs.Debug("show start")
	u := model.NewUser()
	datas, err := u.All()
	if err != nil {
		logs.Error("账号列表获取失败")
		return false
	}

	if len(datas) <= 0 {
		return false
	}
	// 打印数据
	var titlePrinted bool
	var titleSlice []string
	for _, m := range datas {
		// fmt.Printf("%+v\n", m)
		if !titlePrinted {
			rm := reflect.ValueOf(m)
			rm = util.EnsureNotPtrReflectValue(rm)
			for i := 0; i < rm.NumField(); i++ {
				titleSlice = append(titleSlice, rm.Type().Field(i).Name)
			}
			fmt.Printf("| %s |\n", strings.Join(titleSlice, " | "))
			titlePrinted = true
		}
		// fmt.Printf("%v\n", m.(*model.User).GetAge())
		// println(m.(model.User).GetUsername())
		fmt.Printf("| %s |\n", strings.Join(m.ToStringSlice(), " | "))
	}
	return true
}
