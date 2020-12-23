package controller

type controller interface {
	Handle() (bool, nextRoutes)
}

type nextRoutes []routeKey

type controllerNewFunc func() controller

// 只有用到的时候才实例化
type routeCache struct {
	tip      string // 操作提示
	new      controllerNewFunc
	instance controller // 实例
}

func NewRouteCache(tip string, newFunc controllerNewFunc) *routeCache {
	return &routeCache{
		tip: tip,
		new: newFunc,
	}
}

// 执行路由对应的动作
func (cc *routeCache) diaptch() (bool, nextRoutes) {
	if cc.instance == nil {
		cc.instance = cc.new()
	}

	return cc.instance.Handle()
}
