package service

import (
	"gorm.io/gorm"
	"server/global"
	"server/models/program"
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
	//fmt.Println(programs)
	return programs, tx
}
