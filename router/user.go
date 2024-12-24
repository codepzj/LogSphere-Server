package router

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

var userApi = new(v1.UserApi)

func (ur *UserRouter) InitUserRouter(r *gin.Engine) {
	user := r.Group("user")
	{
		user.GET("/get", userApi.GetUserDetailInfo)
		user.GET("/clear", userApi.UserClearStatus)
		user.POST("/create", userApi.UserRegister)
		user.POST("/find", userApi.UserLogin)
		user.POST("/avatar-upload", userApi.UploadUserAvatar)
		user.POST("/edit-profile", userApi.UserEditProfile)
	}
}
