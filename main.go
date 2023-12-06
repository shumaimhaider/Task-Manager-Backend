package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shumaimhaider/task_manager_api/connect"
	"github.com/shumaimhaider/task_manager_api/controller"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	DueDate     time.Time `json:"time"`
}

// Mock Data for tasks
var tasks = []Task{
	{ID: "1", Title: "push", Description: "First task", DueDate: time.Now(), Status: false},
	{ID: "2", Title: "code", Description: "Second task", DueDate: time.Now(), Status: false},
}

func main() {
	db := connect.Connect()
	controller.DB = db

	router := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	router.Use(cors.New(ginConfig))

	router.POST("/create/tasktype", controller.Handle_Create_Task_Type)
	router.PUT("/update/tasktype", controller.Handle_Update_Create_Task_Type)

	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
	})

	router.GET("/task:id", func(ctx *gin.Context) {

		id := ctx.Param("id")
		for _, task := range tasks {
			if task.ID == id {
				ctx.JSON(http.StatusOK, gin.H{
					"task": task,
				})
			}
		}

	})

	router.DELETE("/delete/task/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:i], task)
				ctx.JSON(http.StatusOK, gin.H{
					"Deleted Task": " Task Deleted",
				})

				return
			}

			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Task Not Found",
			})

		}

	})

	router.Run(":8080")
}
