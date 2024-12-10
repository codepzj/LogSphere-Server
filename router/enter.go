package router

type RouterGroup struct {
	UserRouter
	ProgramRouter
}

var RouterGroupApp = new(RouterGroup)
