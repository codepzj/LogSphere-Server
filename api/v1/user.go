package v1

import (
	"fmt"
	"server/models/common/response"
	"server/models/user"
	"server/service"
	"server/utils"

	// "time"

	"github.com/gin-gonic/gin"
)

var userService = new(service.UserService)

type UserApi struct {
}

func (ua *UserApi) UserRegister(c *gin.Context) {
	var u user.UserModel
	if err := c.ShouldBindJSON(&u); err != nil {
		response.FailWithMessage("参数不合法", c)
		return
	}
	fmt.Println(u)

	if err := userService.CreateUser(u); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (ua *UserApi) UserLogin(c *gin.Context) {
	var u user.UserModel
	if err := c.ShouldBindJSON(&u); err != nil {
		response.FailWithMessage("用户参数不合法", c)
		return
	}
	id := userService.GetUserID(u)
	if id == 0 {
		response.FailWithMessage("账号或密码错误", c)
		return
	}
	if cu, isFind := userService.FindUserDetailByID(id); isFind {
		token, _ := utils.GenerateToken(u.Account)
		utils.SetToken(c, token, 86400*7)
		response.OkWithData(cu, c)
		return
	}
	response.FailWithMessage("系统内部错误", c)
}

// UserClearStatus 清空user的Cookie
func (ua *UserApi) UserClearStatus(c *gin.Context) {
	utils.ClearToken(c)
	response.Ok(c)
}

func (ua *UserApi) UploadUserAvatar(c *gin.Context) {
	utils.UploadFile("user", c)
}

func (ua *UserApi) UserEditProfile(c *gin.Context) {
	var ud user.UserDetailModel
	if c.ShouldBindJSON(&ud) != nil {
		response.FailWithMessage("用户参数错误", c)
		return
	}
	fmt.Println(ud)
	if updateRows, err := userService.EditUserDetails(ud); err == nil && updateRows != 0 {
		response.OkWithMessage("更新用户信息成功", c)
		return
	}
	response.FailWithMessage("更新用户信息失败", c)
}
