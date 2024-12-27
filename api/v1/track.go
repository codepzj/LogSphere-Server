package v1

import (
	"fmt"
	"server/models/common/response"
	"server/models/track"
	"server/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

type TrackAPI struct {
}

var (
	trackService = new(service.TrackService)
	regionDBPath = "ip2region.xdb"
)

func (ta *TrackAPI) TrackUser(c *gin.Context) {
	var t track.TrackModel
	if c.ShouldBindJSON(&t) != nil {
		response.FailWithMessage("参数缺少", c)
	}
	clientIP := c.ClientIP()
	t.IPAddr = clientIP
	searcher, err := xdb.NewWithFileOnly(regionDBPath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
	}
	region, err := searcher.SearchByStr(clientIP)
	if err == nil {
		regionSlice := strings.Split(region, "|")
		fmt.Println(regionSlice)
		area := regionSlice[0]
		if area == "0" {
			area = "未知"
		}
		t.Location = area
	}

	if err := trackService.TrackUserAction(t); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)

}

func (ta *TrackAPI) GetAllTrackRecordsByWebsiteId(c *gin.Context) {
	websiteId := c.DefaultQuery("websiteId", "")
	if websiteId == "" {
		response.FailWithMessage("websiteId为空，查询失败", c)
		return
	}
	tm := trackService.GetAllTrackRecordsByWebsiteId(websiteId)
	response.OkWithDetailed(tm, "查询所有记录成功", c)
}

func (ta *TrackAPI) GetAnalyse(c *gin.Context) {
	websiteId := c.DefaultQuery("websiteId", "")
	if websiteId == "" {
		response.FailWithMessage("websiteId为空", c)
		return
	}
	views := trackService.GetAllPageViews(websiteId)
	visitors := trackService.GetVisitorNums(websiteId)
	pageDuration := trackService.GetPageDuration(websiteId)
	referrerInfo := trackService.GetReferrer(websiteId)
	deviceInfo, _ := trackService.GetDeviceInfo(websiteId)
	pageInfo, _ := trackService.GetPageInfo(websiteId)
	locationInfo, _ := trackService.GetLocationInfo(websiteId)

	response.OkWithData(map[string]any{
		"views":        views,
		"visitors":     visitors,
		"pageDuration": pageDuration,
		"referrerInfo": referrerInfo,
		"deviceInfo":   deviceInfo,
		"pageInfo":     pageInfo,
		"locationInfo": locationInfo,
	}, c)
}
