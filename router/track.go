package router

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type TrackRouter struct{}

var trackApi = new(v1.TrackAPI)

func (tr *TrackRouter) InitTrackRouter(r *gin.Engine) {
	tg := r.Group("track")
	{
		tg.GET("/get-all-records", trackApi.GetAllTrackRecordsByWebsiteId)
		tg.GET("/analyse", trackApi.GetAnalyse)
		tg.POST("/", trackApi.TrackUser)
	}
}
