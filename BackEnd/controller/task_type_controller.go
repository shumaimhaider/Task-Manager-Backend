package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shumaimhaider/task_manager_api/models"
)

type TaskTypeRequest struct {
	TaskTypeID   int    `json:"task_type_id"`
	TaskTypeName string `json:"task_type_name"`
}

func HandleCreateTaskType(ctx *gin.Context) {
	var taskTypeRequest TaskTypeRequest
	err := ctx.ShouldBindJSON(&taskTypeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := models.CreateTaskType(ctx, taskTypeRequest.TaskTypeName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error: %w": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func HandleUpdateTaskType(ctx *gin.Context) {
	var taskTypeRequest TaskTypeRequest
	err := ctx.ShouldBindJSON(&taskTypeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := models.UpdatetaskType(ctx, taskTypeRequest.TaskTypeID, taskTypeRequest.TaskTypeName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func HandleGetTaskType(ctx *gin.Context) {
	res, err := models.GetAllTaskType(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func HandleDeleteTaskType(ctx *gin.Context) {
	var taskTypeRequest TaskTypeRequest
	err := ctx.ShouldBindJSON(&taskTypeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := models.DeleteTaskType(ctx, taskTypeRequest.TaskTypeName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
