package track

import (
	"gorm.io/gorm"
	"server/models/program"
)

type TrackModel struct {
	gorm.Model
	Type           string                `json:"type"`      // "pageview" 或 "event"
	VisitorID      string                `json:"visitorId"` // 唯一访客 ID
	URL            string                `json:"url"`       // 页面 URL
	Referrer       string                `json:"referrer"`  // 上一个页面的 URL
	Screen         string                `json:"screen"`    // 屏幕分辨率
	Language       string                `json:"language"`  // 浏览器语言
	UserAgent      string                `json:"userAgent"` // 用户代理信息
	Timestamp      int64                 `json:"timestamp"` // 时间戳
	StayDuration   int64                 `json:"stayDuration"`
	ProgramModelID string                `json:"website_id"`
	ProgramModel   *program.ProgramModel `json:"-" gorm:"references:WebSiteId"`
}

type GraphReq struct {
	WebsiteId string `json:"website_id" binding:"required"`
}
