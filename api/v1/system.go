package v1

import (
	"github.com/gin-gonic/gin"
	"server/models/common/response"
	"server/utils"
)

type SystemApi struct {
}

func (sa *SystemApi) GetSystemUsage(c *gin.Context) {
	usageData := utils.GetSystemUsage()
	response.OkWithData(usageData, c)
}
