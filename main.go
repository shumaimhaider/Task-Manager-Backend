package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/create", func(ctx *gin.Context) {
		var createTask Task
		if err := ctx.ShouldBindJSON(&createTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		tasks = append(tasks, createTask)
		ctx.JSON(http.StatusCreated, gin.H{
			"tasks": tasks,
		})

	})

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

	router.PUT("/update/task/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var updatedTask Task

		if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, task := range tasks {
			if task.ID == id {
				// update only specified field
				if updatedTask.Title != "" {
					tasks[i].Title = updatedTask.Title
				}

				ctx.JSON(http.StatusOK, gin.H{
					"Updated Task": task,
				})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task Not Found",
		})
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
