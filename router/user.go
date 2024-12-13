package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type UserRouter struct {
}

var userApi = new(v1.UserApi)

func (ur *UserRouter) InitUserRouter(r *gin.Engine) {
	user := r.Group("user")
	{
		user.GET("/clear", userApi.UserClearStatus)
		user.POST("/create", userApi.UserRegister)
		user.POST("/find", userApi.UserLogin)
		user.POST("/avatar-upload", userApi.UploadUserAvatar)
	}
}
