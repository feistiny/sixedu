package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/feistiny/sixedu/controller"
	"github.com/feistiny/sixedu/etc"
	"github.com/feistiny/sixedu/util"
	"log"
)

func init() {
	var err error
	err = etc.LoadConfig("etc/config.json")
	if err != nil {
		log.Fatalln("配置文件读取失败", err)
	}
	initLog()
}

func initLog() {
	var err error
	ll, err := etc.Cfg.Int("LOG_LEVEL")
	if err != nil {
		log.Fatalln("日志级别读取错误", err)
	}
	// 级别如下:
	// const (
	//	LevelEmergency = iota
	//	LevelAlert
	//	LevelCritical
	//	LevelError
	//	LevelWarning
	//	LevelNotice
	//	LevelInformational
	//	LevelDebug
	// )
	logs.SetLevel(ll)
}

func main() {
	var logged bool
	util.AutoQuit()
	util.TipStart()
labelStart:
	if !logged {
		println("欢迎来到 六星教育学员管理系统 sixedu")
		for {
			println("你的执行操作:")
			util.OpTips([]string{"登入系统", "注册用户"})
			choice := util.GetInput()
			util.TipEnd()
			println()
			util.TipStart()
			switch choice {
			case "0":
				c := controller.NewLogin()
				if c.Handle() {
					logged = true
					println("登入成功")
					goto labelStart
				} else {
					println("登入失败")
				}
			case "1":
				c := controller.NewReg()
				if c.Handle() {
					println("注册成功")
				} else {
					println("注册失败")
				}
			default:
				println("请重新输入")
			}
		}
	} else {
		util.OpTips([]string{"展示用户信息"})
		choice := util.GetInput()
		util.TipEnd()
		println()
		util.TipStart()
		switch choice {
		case "0":
			c := controller.NewShow()
			if c.Handle() {
				logged = true
				println("展示成功")
			} else {
				println("展示失败")
			}
		default:
			println("请重新输入")
		}
	}
	util.TipEnd()
}
