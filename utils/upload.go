package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"server/models/common/response"
)

// UploadFile dir指的是上传路径为/uploads/<dir>/demo.jpg|.png.|webp
func UploadFile(dir string, c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("文件上传失败", c)
		return
	}

	fileExt := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%s%s", uuid.NewString(), fileExt)

	dst := filepath.Join("uploads", dir, newFileName)

	// 创建上传目录，如果没有则创建
	baseDir := "uploads" + "/" + dir
	err = os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		response.FailWithMessage("无法创建文件夹", c)
		return
	}

	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		response.FailWithMessage("文件无法保存", c)
		return
	}
	response.OkWithDetailed(map[string]any{"filename": newFileName}, "文件上传成功", c)
}
