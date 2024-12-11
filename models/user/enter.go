package user

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Account  string `json:"account,omitempty" gorm:"unique"`
	Password string `json:"password"`
}

type UserDetailModel struct {
	Nickname    string
	Role        int
	Avatar      string
	UserModelID uint      `json:"account_id"`
	UserModel   UserModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
