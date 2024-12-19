package router

type RouterGroup struct {
	UserRouter
	ProgramRouter
	TrackRouter
}

var RouterGroupApp = new(RouterGroup)
