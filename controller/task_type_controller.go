package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shumaimhaider/task_manager_api/models"
	"github.com/uptrace/bun"
)

var DB *bun.DB

type TaskTypeRequest struct {
	TaskTypeID   int    `json:"task_type_id"`
	TaskTypeName string `json:"task_type_name"`
}

func Handle_Create_Task_Type(ctx *gin.Context) {
	var taskTypeRequest TaskTypeRequest
	err := ctx.ShouldBindJSON(&taskTypeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	res, err := models.Create_Task_Type(ctx, taskTypeRequest.TaskTypeName, DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error creating task type: %w": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, res)
}

func Handle_Update_Create_Task_Type(ctx *gin.Context) {
	var taskTypeRequest TaskTypeRequest
	err := ctx.ShouldBindJSON(&taskTypeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	res, err := models.Update_task_Type(ctx, taskTypeRequest.TaskTypeID, taskTypeRequest.TaskTypeName, DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, res)
}
