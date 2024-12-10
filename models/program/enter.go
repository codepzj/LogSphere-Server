package program

import "gorm.io/gorm"

type ProgramModel struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	Domain    string `json:"domain" binding:"required"`
	Secure    bool   `json:"secure" binding:"required"`
	AccountID uint   `json:"account_id" gorm:"column:account_id"`
	WebSiteId string `json:"-" gorm:"column:website_id"`
}
