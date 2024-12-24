package user

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Account  string `json:"account,omitempty" gorm:"unique"`
	Password string `json:"password"`
}

type UserDetailModel struct {
	Nickname    string    `json:"nickname,omitempty"`
	Role        int       `json:"role"`
	Avatar      string    `json:"avatar"`
	UserModelID uint      `json:"account_id,omitempty"`
	UserModel   UserModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_model"`
}
