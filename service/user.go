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
	err := db.First(&user.UserModel{}, "account = ?", u.Account).Error
	if err == nil {
		return errors.New("账号已存在，无法重复创建")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("注册失败，请咨询网站管理员")
	}
	userDetail := user.UserDetailModel{
		Nickname:  "未命名",
		Role:      0,
		Avatar:    "",
		UserModel: u,
	}
	if db.Create(&userDetail).Error != nil {
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

func (us *UserService) GetUserID(u user.UserModel) uint {
	var cu user.UserModel
	if global.LS_DB.Where("account = ? and password = ?", u.Account, u.Password).First(&cu).Error == nil {
		return cu.ID
	}
	return 0
}

// FindUserDetailByID 通过id查找user是否存在并返回详细信息
func (us *UserService) FindUserDetailByID(id uint) (user.UserDetailModel, bool) {
	// 获取完整的user信息
	var cu user.UserDetailModel
	if global.LS_DB.Preload("UserModel").First(&cu, user.UserDetailModel{UserModelID: id}).Error == nil {
		return cu, true
	}
	return user.UserDetailModel{}, false
}

func (us *UserService) EditUserDetails(u user.UserDetailModel) (int, error) {
	result := global.LS_DB.Model(&user.UserDetailModel{}).Where("user_model_id = ?", u.UserModelID).Select("*").Updates(u)
	return int(result.RowsAffected), result.Error
}

func (us *UserService) GetAllUsers() []user.UserDetailModel {
	var users []user.UserDetailModel
	global.LS_DB.Preload("UserModel").Find(&users)
	return users
}
