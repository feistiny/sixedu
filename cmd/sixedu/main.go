package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/feistiny/sixedu/controller"
	"github.com/feistiny/sixedu/etc"
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
	controller.Run()
}
