package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/global"
)

func ConnectDataBase() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:pzj20162116@tcp(127.0.0.1:3306)/logsphere?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         191,                                                                                    // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                                                                                  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	global.LS_DB = db
}
