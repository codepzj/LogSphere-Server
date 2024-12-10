package service

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/models/user"
)

type UserService struct{}

// CreateUser 新增用户
func (us *UserService) CreateUser(u user.UserModel) error {
	db := global.LS_DB
	err := db.First(&user.UserModel{}, "account=?", u.Account).Error
	if err == nil {
		return errors.New("账号已存在，无法重复创建")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("注册失败，请咨询网站管理员")
	}
	if db.Create(&u).Error != nil {
		return errors.New("注册失败，请重试")
	}
	return nil
}

// FindUser 查找user是否存在
func (us *UserService) FindUser(u user.UserModel) (user.UserModel, bool) {
	// 获取完整的user信息
	var cu user.UserModel
	if global.LS_DB.Where(&u).First(&cu).Error == nil {
		return cu, true
	}
	return user.UserModel{}, false
}

func (us *UserService) EditUserDetails(id string, u user.UserDetailModel) {

}
