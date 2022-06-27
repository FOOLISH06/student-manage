package main

import (
	"github.com/gin-gonic/gin"
	"student-manage/controller"
)

func collectRouter(router *gin.Engine) *gin.Engine {
	root := router.Group("/studentMS")

	// 管理者api
	manager := root.Group("/api/manager")
	manager.POST("/register", controller.Register)

	// 学生api
	//student := root.Group("/api/student")

	return router
}
