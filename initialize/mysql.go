package initialize

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"server/global"
)

func ConnectDataBase() {
	// 加载 .env 文件中的环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// 从环境变量中获取数据库连接配置
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbLoc := os.Getenv("DB_LOC")

	// 构建 DSN 数据源名称
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbLoc)

	// 使用 GORM 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // 使用从环境变量中读取的 DSN
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	global.LS_DB = db
}
