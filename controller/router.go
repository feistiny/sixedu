package controller

import (
	"github.com/feistiny/sixedu/util"
)

var CurrentRoute = IndexRoute

type routeKey int

const (
	IndexRoute routeKey = iota
	LoginRoute
	RegRoute
	ShowRoute
)

var routePool = make(map[routeKey]*routeCache)

var cachedNextKeys []routeKey

func init() {
	routePool[IndexRoute] = NewRouteCache("欢迎首页", NewIndex)
	routePool[LoginRoute] = NewRouteCache("登入账号", NewLogin)
	routePool[RegRoute] = NewRouteCache("注册信息", NewReg)
	routePool[ShowRoute] = NewRouteCache("展示所有账号", NewShow)

	// 按 x 后回车自动退出, 不用每个地方都判断
	util.AutoQuit()
}

func Run() {
	for {
		util.TipStart()
		dispatch(CurrentRoute)
		util.TipEnd()
		println()
	}
}

func dispatch(key routeKey) {
	rc := getRouteCache(key)
	success, nextKeys := rc.diaptch()
	if !success {
		if len(nextKeys) == 0 {
			nextKeys = cachedNextKeys
		}
		println(rc.tip + "失败")
	} else {
		println(rc.tip + "成功")
	}
	if len(nextKeys) > 0 {
		cachedNextKeys = nextKeys
	}

	// 提示选择下一个操作
	tips, choices := getRouteTips(nextKeys)
	util.OpTips(tips)
	for {
		choice, err := util.GetChoice("你的执行操作")
		if err != nil {
			println(err)
			return
		}
		if choice >= len(choices) {
			println("你的选择不存在, 请重新选择")
		}

		CurrentRoute = nextKeys[choice]
		break
	}
}

func getRouteTips(keys []routeKey) (tips []string, maps []*routeCache) {
	for _, key := range keys {
		tips = append(tips, getRouteCache(key).tip)
		maps = append(maps, getRouteCache(key))
	}
	return
}

func getRouteCache(key routeKey) (rc *routeCache) {
	if _, ok := routePool[key]; !ok {
		println("无效路由, 请重新输入")
	}
	rc = routePool[key]
	return
}
