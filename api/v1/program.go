package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/models/common/response"
	"server/models/program"
	"server/service"
)

var programService = new(service.ProgramService)

type ProgramApi struct{}

func (pa *ProgramApi) ProgramCreate(c *gin.Context) {
	var newProgram program.ProgramModel
	if c.ShouldBindJSON(&newProgram) != nil {
		response.Fail(c)
		return
	}
	userId := uuid.NewString()
	newProgram.WebSiteId = userId
	if programService.ProgramCreate(newProgram) != nil {
		response.FailWithMessage("创建项目失败", c)
	}
	response.Ok(c)
}

func (pa *ProgramApi) ProgramFindAll(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithDetailed(nil, "参数缺失", c)
		return
	}
	if programs, tx := programService.FindAllProgramByAccountID(id); tx.Error == nil {
		response.OkWithData(programs, c)
		return
	}
	response.FailWithMessage("查询失败", c)
}
