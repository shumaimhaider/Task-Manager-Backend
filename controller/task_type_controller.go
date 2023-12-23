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
	}
	res, err := models.CreateTaskType(ctx, taskTypeRequest.TaskTypeName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error creating task type: %w": err.Error(),
		})
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
	}

	res, err := models.UpdatetaskType(ctx, taskTypeRequest.TaskTypeID, taskTypeRequest.TaskTypeName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, res)
}

func HandleGetTaskType(ctx *gin.Context) {
	res, err := models.GetAllTaskType(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
	}

	res, err := models.DeleteTaskTypeType(ctx, taskTypeRequest.TaskTypeName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, res)
}
