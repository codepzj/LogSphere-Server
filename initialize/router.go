package initialize

import (
	"server/middleware"
	"server/router"

	"github.com/gin-gonic/gin"
)

var (
	UserRouter    = router.RouterGroupApp.UserRouter
	ProgramRouter = router.RouterGroupApp.ProgramRouter
	TrackRouter   = router.RouterGroupApp.TrackRouter
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Static("/uploads", "./uploads")
	UserRouter.InitUserRouter(r)
	ProgramRouter.InitProgramRouter(r)
	TrackRouter.InitTrackRouter(r)
	return r
}
