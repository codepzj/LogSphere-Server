package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type TrackRouter struct{}

var trackApi = new(v1.TrackAPI)

func (tr *TrackRouter) InitTrackRouter(r *gin.Engine) {
	tg := r.Group("track")
	{
		tg.GET("/get-all-records", trackApi.GetAllTrackRecordsByWebsiteId)
		tg.GET("/", trackApi.TrackUser)
		tg.POST("/analyse", trackApi.GetAnalyse)
	}
}
