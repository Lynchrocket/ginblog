package main

import (
	"ginblog/models"
	"ginblog/routes"
)

func main() {
	// 引用数据库
	models.InitDb()

	routes.InitRouter()
}
