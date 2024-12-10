package user

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Account  string `json:"account,omitempty" gorm:"unique"`
	Password string `json:"password,omitempty"`
}

type UserDetailModel struct {
	Nickname    string
	Role        int
	Avatar      string
	UserModelID uint
	UserModel   UserModel
}
