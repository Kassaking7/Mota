package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoController struct{}

func NewInfoController() *InfoController {
	return &InfoController{}
}

func (c *InfoController) GetInfo(ctx *gin.Context) {
	// 创建一个map作为JSON响应的内容
	info := map[string]string{"message": "Hello, World!"}

	// 返回JSON响应
	ctx.JSON(http.StatusOK, info)
}
func ImportRoutes(r *gin.Engine) {
	// 创建控制器实例
	infoController := NewInfoController()

	// 定义路由
	r.GET("/info", infoController.GetInfo)
}
