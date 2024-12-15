package program

import (
	"gorm.io/gorm"
)

type ProgramModel struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	Domain    string `json:"domain" binding:"required"`
	Secure    bool   `json:"secure"`
	AccountID uint   `json:"account_id" binding:"required" gorm:"column:account_id"`
	WebSiteId string `json:"website_id,omitempty" gorm:"column:website_id;unique"`
}
