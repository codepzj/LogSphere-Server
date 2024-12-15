package initialize

import (
	"server/global"
	"server/models/program"
	"server/models/track"
	"server/models/user"
)

func RegisterTables() {
	err := global.LS_DB.AutoMigrate(
		&user.UserModel{},
		&user.UserDetailModel{},
		&program.ProgramModel{},
		&track.TrackModel{},
	)
	if err != nil {
		panic("迁移数据库失败")
	}
}
