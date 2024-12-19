package service

import (
	"server/global"
	"server/models/track"
)

type TrackService struct {
}

// TrackUserAction 只有在页面进入和离开的时候将追踪日志写入数据库
func (ts *TrackService) TrackUserAction(tm track.TrackModel) error {
	switch tm.Type {
	case "pageview", "pageStayTime":
		return global.LS_DB.Create(&tm).Error

	case "heartbeat":
		return nil
	default:
		return nil
	}

}

func (ts *TrackService) GetAllTrackRecordsByWebsiteId(websiteId string) []track.TrackModel {
	var tm []track.TrackModel
	global.LS_DB.Order("id desc").Find(&tm, track.TrackModel{ProgramModelID: websiteId})
	return tm
}

// GetAllPageViews 页面访问数
func (ts *TrackService) GetAllPageViews(websiteId string) int64 {
	var count int64
	global.LS_DB.Model(&track.TrackModel{}).Where(track.TrackModel{ProgramModelID: websiteId}).Where("type = ?", "pageview").Count(&count)
	return count
}

func (ts *TrackService) GetVisitorNums(websiteId string) int64 {
	var count int64
	global.LS_DB.Model(&track.TrackModel{}).Where(track.TrackModel{ProgramModelID: websiteId}).Distinct("visitor_id").Count(&count)
	return count
}

func (ts *TrackService) GetActiveUsersNum(websiteId string) int64 {
	var count int64
	global.LS_DB.Model(&track.TrackModel{}).Where(track.TrackModel{ProgramModelID: websiteId}).Where("type = ?", "heart").Count(&count)
	return count
}
