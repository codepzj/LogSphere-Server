package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net/http"
	"server/models/common/response"
	"server/models/track"
	"server/service"
	"strings"
)

type TrackAPI struct {
}

var (
	trackService = new(service.TrackService)
	regionDBPath = "ip2region.xdb"
)

// WebSocket upgrader configuration
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin for simplicity
	},
}

// TrackUser handles incoming WebSocket messages (tracking user actions)
func (ta *TrackAPI) TrackUser(c *gin.Context) {
	// 升级到 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}()

	for {
		// 判断连接是否已关闭
		if conn == nil {
			fmt.Println("WebSocket connection is closed")
			return
		}

		// Read message from WebSocket connection
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				fmt.Println("Connection closed by client.")
			} else {
				fmt.Println("Error reading message:", err)
			}
			return
		}

		// 打印收到的消息内容
		fmt.Println("Received message:", string(msg)) // 打印接收到的消息

		var t track.TrackModel
		if err := json.Unmarshal(msg, &t); err != nil {
			fmt.Println("Error unmarshalling message:", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid data format"))
			return
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
			fmt.Println("Error tracking user action:", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Failed to create tracking record"))
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte("Tracking successful")); err != nil {
			fmt.Println("Error sending acknowledgment:", err)
			return
		}
	}
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
