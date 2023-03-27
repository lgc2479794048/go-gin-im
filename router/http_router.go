/*
 * @Author: lgc2479794048 lgc2479794048@gmail.com
 * @Date: 2023-03-25 15:04:47
 * @LastEditors: lgc2479794048 lgc2479794048@gmail.com
 * @LastEditTime: 2023-03-26 14:07:29
 * @FilePath: \go-gin-im\router\http_router.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package router

import (
	"fmt"
	"go-gin-im/config"
	"net/http"
	_ "net/http/pprof"

	uploadController "go-gin-im/controller/upload"
	userController "go-gin-im/controller/user"

	"github.com/gin-gonic/gin"
)

func AppStart() {
	appConfig, err := config.NewAppConfig()
	if err != nil {
		// 处理错误
		panic(err)
	}
	// 设置gin框架的模式
	if appConfig.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := InitRouter()
	// r.Use(middleware.ErrorHandler())
	r.Run(fmt.Sprintf(":%d", appConfig.Server.Port)) // 启动服务，并监听 8080 端口
}

func InitRouter() *gin.Engine {
	// 创建一个 Gin 实例
	router := gin.Default()
	// 注册 pprof 处理程序
	router.GET("/debug/pprof/*pprof", func(c *gin.Context) {
		http.DefaultServeMux.ServeHTTP(c.Writer, c.Request)
	})
	// 定义路由组
	api := router.Group("/go-gin-im/user")
	{
		// 定义接口路由
		api.POST("/register", userController.Register) // 注册接口
		api.GET("/login", userController.Login)        // 登录接口
	}
	uploadApi := router.Group("/go-gin-im/upload")
	{
		uploadApi.POST("/uploadFile", uploadController.Upload) // 上传文件接口
	}

	// 返回 Gin 实例
	return router
}
