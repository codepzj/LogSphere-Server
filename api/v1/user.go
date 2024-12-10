package v1

import (
	"server/models/common/response"
	"server/models/user"
	"server/service"
	"server/utils"

	// "time"

	"github.com/gin-gonic/gin"
)

var usrService = new(service.UserService)

type UserApi struct {
}

func (ua *UserApi) UserRegister(c *gin.Context) {
	var u user.UserModel
	if err := c.ShouldBindJSON(&u); err != nil {
		response.FailWithMessage("参数不合法", c)
		return
	}

	if err := usrService.CreateUser(u); err != nil {
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

	if cu, isFind := usrService.FindUser(u); isFind {
		token, _ := utils.GenerateToken(u.Account)
		utils.SetToken(c, token, 86400*7)
		response.OkWithData(map[string]any{"account_id": cu.ID, "account": cu.Account}, c)
		return
	}
	response.FailWithMessage("账号或密码错误", c)
}

// UserClearStatus 清空user的Cookie
func (ua *UserApi) UserClearStatus(c *gin.Context) {
	utils.ClearToken(c)
	response.Ok(c)
}