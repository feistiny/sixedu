package router

import (
	"github.com/feistiny/sixedu/controller"
	"github.com/feistiny/sixedu/util"
)

var CurrentRoute = controller.IndexRoute

var routePool = make(map[controller.RouteKey]*routeCache)

var cachedNextKeys []controller.RouteKey

func init() {
	routePool[controller.IndexRoute] = newRouteCache("欢迎首页", controller.NewIndex)
	routePool[controller.LoginRoute] = newRouteCache("登入账号", controller.NewLogin)
	routePool[controller.RegRoute] = newRouteCache("注册信息", controller.NewReg)
	routePool[controller.ShowRoute] = newRouteCache("展示所有账号", controller.NewShow)

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

func dispatch(key controller.RouteKey) {
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

func getRouteTips(keys []controller.RouteKey) (tips []string, maps []*routeCache) {
	for _, key := range keys {
		tips = append(tips, getRouteCache(key).tip)
		maps = append(maps, getRouteCache(key))
	}
	return
}

func getRouteCache(key controller.RouteKey) (rc *routeCache) {
	if _, ok := routePool[key]; !ok {
		println("无效路由, 请重新输入")
	}
	rc = routePool[key]
	return
}
