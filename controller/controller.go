package controller

type controller interface {
	Handle() bool
}