package main

import (
	"student-manage/controller"
	"student-manage/middleware"

	"github.com/gin-gonic/gin"
)

func collectRouter(router *gin.Engine) *gin.Engine {
	router.Use(middleware.Cors())
	root := router.Group("/studentMS")

	// 管理者api
	manager := root.Group("/api/manager")
	manager.POST("/login", controller.Login)
	manager.POST("/register", middleware.AuthSuperManagerMid(), controller.Register)

	// 学生api
	student := root.Group("/api/student")
	student.Use(middleware.AuthMiddleware())                    // 添加token验证中间件
	student.GET("/", controller.GetStudents)                    // 查询所有学生
	student.GET("/:sid", controller.GetStudentById)             // 通过学号查询学生
	student.GET("/class/:class", controller.GetStudentsByClass) // 通过班级查询学生
	student.POST("/", controller.CreateStudent)                 // 添加学生
	student.PUT("/", controller.UpdateStudent)                  // 通过学号修改学生信息
	student.DELETE("/:sid", controller.DeleteStudent)           // 通过学号删除学生

	// 访客（学生，只能查询所有学生信息）
	root.POST("/student/login", controller.StudentLogin)
	root.GET("/student/get", controller.StudentGet)
	return router
}
