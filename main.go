package main

import (
	"gin-admin/database"
	"gin-admin/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := router.GetRouter()
	// 使用环境变量中的值
	appPort := os.Getenv("APP_PORT")
	// 监听并在 0.0.0.0:9010 上启动服务
	router.Run(":" + appPort)

	database.CloseMysql()
}
