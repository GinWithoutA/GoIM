package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 设置路由
	router := gin.Default()

	// 设置路由对应 Handler，也就是 Controller
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": "pong",
		})
	})

	// 启动监听，默认监听 8080 端口
	router.Run()
}
