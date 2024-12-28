package router

type RouterGroup struct {
	UserRouter
	ProgramRouter
	TrackRouter
	SystemRouter
}

var RouterGroupApp = new(RouterGroup)
