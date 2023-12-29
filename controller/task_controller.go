package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shumaimhaider/task_manager_api/common"
	"github.com/shumaimhaider/task_manager_api/models"
)

func HandleCreateTask(ctx *gin.Context) {
	var taskRequest common.TaskRequest
	err := ctx.ShouldBindJSON(&taskRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := models.CreateTask(ctx, taskRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error: %w": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func HandleUpdateTask(ctx *gin.Context) {
	var taskRequest common.TaskRequest
	err := ctx.ShouldBindJSON(&taskRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := models.UpdateTaskCompleted(ctx, taskRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func HandleGetTask(ctx *gin.Context) {
	res, err := models.GetAllTask(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func HandleDeleteTask(ctx *gin.Context) {
	var taskRequest common.TaskRequest
	err := ctx.ShouldBindJSON(&taskRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := models.DeleteTask(ctx, taskRequest.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
