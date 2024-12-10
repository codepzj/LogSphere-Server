package core

import (
	"server/initialize"
)

func RunServer() {
	r := initialize.InitRouter()
	err := r.Run(":8080")
	if err != nil {
		panic("服务启动失败")
	}
}
