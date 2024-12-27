package service

import (
	"errors"
	"fmt"
	"server/global"
	"server/models/program"
	"server/models/track"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type TrackService struct {
}

// TrackUserAction 只有在页面进入和离开的时候将追踪日志写入数据库
func (ts *TrackService) TrackUserAction(tm track.TrackModel) error {
	fmt.Println(tm.ProgramModelID)
	err := global.LS_DB.Where("website_id = ?", tm.ProgramModelID).First(&program.ProgramModel{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("项目未找到")
	}
	switch tm.Type {
	case "pageview", "pageduration", "pagebounce":
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

// GetAllPageViews 获取页面访问数
func (ts *TrackService) GetAllPageViews(websiteId string) int64 {
	var count int64
	global.LS_DB.Model(&track.TrackModel{}).Where(track.TrackModel{ProgramModelID: websiteId}).Where("type = ?", "pageview").Count(&count)
	return count
}

// GetVisitorNums 获取游客数量
func (ts *TrackService) GetVisitorNums(websiteId string) int64 {
	var count int64
	global.LS_DB.Model(&track.TrackModel{}).Where(track.TrackModel{ProgramModelID: websiteId}).Where("type = ?", "pageview").Distinct("visitor_id").Count(&count)
	return count
}

func (ts *TrackService) GetActiveUsersNum(websiteId string) int64 {
	var count int64
	global.LS_DB.Model(&track.TrackModel{}).Where(track.TrackModel{ProgramModelID: websiteId}).Where("type = ?", "heartbeat").Count(&count)
	return count
}

// GetPageDuration 获取平均停留时长
func (ts *TrackService) GetPageDuration(websiteId string) float64 {
	var stayDuration []int64
	global.LS_DB.Model(&track.TrackModel{}).Where(track.TrackModel{ProgramModelID: websiteId}).Where("type = ?", "pageduration").Pluck("stay_duration", &stayDuration)
	if len(stayDuration) == 0 {
		return 0
	}
	return funk.Sum(stayDuration) / 1000.0 / float64(len(stayDuration))
}

func (ts *TrackService) GetReferrer(websiteId string) []track.ReferrerRatioModel {
	var referrerList []string
	global.LS_DB.Model(&track.TrackModel{}).
		Where(track.TrackModel{ProgramModelID: websiteId}).
		Where("type = ?", "pageview").
		Pluck("referrer", &referrerList)

	if len(referrerList) == 0 {
		return nil
	}

	// 获取referer中来源以及对应比例
	var referrerRatio []track.ReferrerRatioModel
	uniqueReferrerList := funk.Uniq(referrerList).([]string)
	totalReferrerCount := len(referrerList)

	for _, domain := range uniqueReferrerList {
		filtered := funk.Filter(referrerList, func(x string) bool {
			return x == domain
		}).([]string)

		ratio := float64(len(filtered)) / float64(totalReferrerCount)
		referrerRatio = append(referrerRatio, track.ReferrerRatioModel{
			Domain: domain,
			Count:  len(filtered),
			Ratio:  ratio,
		})
	}

	return referrerRatio
}

func (ts *TrackService) GetDeviceInfo(websiteId string) (map[string]any, error) {
	var screen []string
	// Fetch screen data from the database for the given website
	err := global.LS_DB.Model(&track.TrackModel{}).
		Where(track.TrackModel{ProgramModelID: websiteId}).
		Where("type = ?", "pageview").
		Pluck("screen", &screen).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch screen data: %v", err)
	}

	// Initialize device counts
	deviceCounts := map[string]int{"mobile": 0, "laptop": 0, "desktop": 0}
	validCount := 0

	// Function to categorize device based on width
	categorizeDevice := func(width int) string {
		switch {
		case width < 768:
			return "mobile"
		case width <= 1200:
			return "laptop"
		default:
			return "desktop"
		}
	}

	// Process each screen resolution and categorize devices
	for _, s := range screen {
		if len(s) == 0 || !strings.Contains(s, "x") {
			continue
		}

		dimensions := strings.Split(s, "x")
		if len(dimensions) != 2 {
			continue
		}

		width, err := strconv.Atoi(dimensions[0])
		if err != nil {
			continue // Skip invalid width
		}

		// Increment the appropriate device count
		deviceType := categorizeDevice(width)
		deviceCounts[deviceType]++
		validCount++
	}

	// If no valid screen data is available, return an error
	if validCount == 0 {
		return nil, fmt.Errorf("no valid screen data available")
	}

	// Calculate the total count of valid devices
	totalCount := float64(validCount)

	// Construct the result map with counts and ratios
	result := map[string]any{
		"mobile": map[string]any{
			"count": deviceCounts["mobile"],
			"ratio": float64(deviceCounts["mobile"]) / totalCount,
		},
		"laptop": map[string]any{
			"count": deviceCounts["laptop"],
			"ratio": float64(deviceCounts["laptop"]) / totalCount,
		},
		"desktop": map[string]any{
			"count": deviceCounts["desktop"],
			"ratio": float64(deviceCounts["desktop"]) / totalCount,
		},
	}

	return result, nil
}

func (ts *TrackService) GetPageInfo(websiteId string) (map[string]map[string]float64, error) {
	// Fetch all the page URLs for the specified website
	var pageList []string
	err := global.LS_DB.Model(&track.TrackModel{}).
		Where(track.TrackModel{ProgramModelID: websiteId}).
		Where("type = ?", "pageview").
		Pluck("url", &pageList).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch page URLs: %v", err)
	}

	// Check if there are no page views
	if len(pageList) == 0 {
		return nil, fmt.Errorf("no pageview data found for the website")
	}

	// Count the total number of page views
	totalPageViews := float64(len(pageList))

	// Initialize a map to count occurrences of each unique URL
	urlCounts := make(map[string]int)
	for _, url := range pageList {
		urlCounts[url]++
	}

	// Create a result map to store both count and ratio for each URL
	result := make(map[string]map[string]float64)
	for url, count := range urlCounts {
		result[url] = map[string]float64{
			"count": float64(count),
			"ratio": float64(count) / totalPageViews,
		}
	}

	return result, nil
}

func (ts *TrackService) GetLocationInfo(websiteId string) (map[string]map[string]float64, error) {
	// Fetch all the location data for the specified website
	var locationList []string
	err := global.LS_DB.Model(&track.TrackModel{}).
		Where(track.TrackModel{ProgramModelID: websiteId}).
		Where("type = ?", "pageview").
		Pluck("location", &locationList).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch location data: %v", err)
	}

	// Check if there are no location data
	if len(locationList) == 0 {
		return nil, fmt.Errorf("no location data found for the website")
	}

	// Count the total number of page views
	totalPageViews := float64(len(locationList))

	// Initialize a map to count occurrences of each unique location
	locationCounts := make(map[string]int)
	for _, location := range locationList {
		locationCounts[location]++
	}

	// Create a result map to store both count and ratio for each location
	result := make(map[string]map[string]float64)
	for location, count := range locationCounts {
		result[location] = map[string]float64{
			"count": float64(count),
			"ratio": float64(count) / totalPageViews,
		}
	}

	return result, nil
}
