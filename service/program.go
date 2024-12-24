package service

import (
	"gorm.io/gorm"
	"server/global"
	"server/models/program"
	"server/models/track"
)

type ProgramService struct {
}

// ProgramCreate 增加项目
func (ps *ProgramService) ProgramCreate(pg program.ProgramModel) error {
	return global.LS_DB.Create(&pg).Error
}

func (ps *ProgramService) FindAllProgramByAccountID(AccountID string) ([]program.ProgramModel, *gorm.DB) {
	var programs []program.ProgramModel
	tx := global.LS_DB.Where("account_id = ?", AccountID).Find(&programs)
	return programs, tx
	
}
func (ps *ProgramService) FindURLByWebsiteID(websiteId string) string {
	var p program.ProgramModel
	global.LS_DB.Where("website_id = ?", websiteId).First(&p)
	if p.Secure {
		return "https://" + p.Domain
	}
	
	return "http://" + p.Domain
}

func (ps *ProgramService) DeleteProgram(websiteId string) error {
	if err := global.LS_DB.Where("website_id = ?", websiteId).Delete(&track.TrackModel{}).Error; err != nil {
		return err
	}
	// 同时删除父模型
	return global.LS_DB.Where("website_id = ?", websiteId).Delete(&program.ProgramModel{}).Error
}
