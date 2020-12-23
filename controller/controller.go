package controller

type Controller interface {
	Handle() (bool, NextRoutes)
}

type NextRoutes []RouteKey

type RouteKey int

const (
	IndexRoute RouteKey = iota
	LoginRoute
	RegRoute
	ShowRoute
)
