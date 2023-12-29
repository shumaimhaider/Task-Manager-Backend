package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shumaimhaider/task_manager_api/common"
	"github.com/shumaimhaider/task_manager_api/connect"
	"github.com/shumaimhaider/task_manager_api/controller"
)

func main() {
	db := connect.Connect()
	common.DB = db

	router := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	router.Use(cors.New(ginConfig))

	router.POST("/create/tasktype", controller.HandleCreateTaskType)
	router.PUT("/update/tasktype", controller.HandleUpdateTaskType)
	router.GET("/taskstype", controller.HandleGetTaskType)
	router.DELETE("/delete/tasktype", controller.HandleDeleteTaskType)

	router.POST("/create/task", controller.HandleCreateTask)
	router.PUT("/update/task", controller.HandleUpdateTask)
	router.GET("/tasks", controller.HandleGetTask)
	router.DELETE("/delete/task", controller.HandleDeleteTask)

	router.Run(":8080")
}
