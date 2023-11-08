package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 导入路由
	ImportRoutes(r)

	r.Run(":8088")
}
