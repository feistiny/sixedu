package router

import "github.com/feistiny/sixedu/controller"

type controllerNewFunc func() controller.Controller

// 只有用到的时候才实例化
type routeCache struct {
	tip      string // 操作提示
	new      controllerNewFunc
	instance controller.Controller // 实例
}

func newRouteCache(tip string, newFunc controllerNewFunc) *routeCache {
	return &routeCache{
		tip: tip,
		new: newFunc,
	}
}

// 执行路由对应的动作
func (cc *routeCache) diaptch() (bool, controller.NextRoutes) {
	if cc.instance == nil {
		cc.instance = cc.new()
	}

	return cc.instance.Handle()
}
