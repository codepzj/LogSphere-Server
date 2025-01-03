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
	newProgram.WebSiteId = uuid.NewString()
	if programService.ProgramCreate(newProgram) != nil {
		response.FailWithMessage("创建项目失败", c)
	}
	response.OkWithData(map[string]string{
		"website_id": newProgram.WebSiteId,
	}, c)
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

func (pa *ProgramApi) URLFindByWebsiteId(c *gin.Context) {
	websiteId := c.DefaultQuery("websiteId", "")
	if websiteId == "" {
		response.FailWithDetailed(nil, "参数缺失", c)
		return
	}
	url := programService.FindURLByWebsiteID(websiteId)
	response.OkWithData(map[string]string{
		"url": url,
	}, c)
}
func (pa *ProgramApi) DeleteProgram(c *gin.Context) {
	websiteId := c.Param("websiteId")
	if websiteId == "" {
		response.FailWithMessage("website_id为空，删除失败", c)
		return
	}
	if err := programService.DeleteProgram(websiteId); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
