package controller

import "github.com/feistiny/sixedu/router"

type Controller interface {
	Handle() (bool, router.NextRoutes)
}
