package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

type SystemRouter struct{}

var SystemApi = new(v1.SystemApi)

func (sr *SystemRouter) InitSystemRouter(r *gin.Engine) {
	sg := r.Group("system").Use(middleware.JWTAuth())
	{
		sg.GET("/usage", SystemApi.GetSystemUsage)
	}
}
