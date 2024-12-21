package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

type ProgramRouter struct{}

var ProgramApi = new(v1.ProgramApi)

func (pr *ProgramRouter) InitProgramRouter(r *gin.Engine) {
	pg := r.Group("program").Use(middleware.JWTAuth())
	{
		pg.GET("/find/:id", ProgramApi.ProgramFindAll)
		pg.GET("/find/domain-by-websiteId", ProgramApi.URLFindByWebsiteId)
		pg.POST("/create", ProgramApi.ProgramCreate)
	}
}
