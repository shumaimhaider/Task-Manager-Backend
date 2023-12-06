package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type TaskType struct {
	bun.BaseModel `bun:"table:task_type"`
	TaskTypeId    int    `bun:"task_type_id,pk,autoincrement"`
	TaskTypeName  string `bun:"task_type_name,notempty"`
}

func Create_Task_Type(ctx *gin.Context, task_type_name string, db *bun.DB) (int, error) {
	taskType := &TaskType{
		TaskTypeName: task_type_name,
	}

	_, err := db.NewInsert().Model(taskType).Exec(ctx)
	if err != nil {
		return -1, fmt.Errorf("error inserting task type: %w", err)
	}

	return taskType.TaskTypeId, nil
}

func Update_task_Type(ctx *gin.Context, task_type_id int, task_type_name string, db *bun.DB) (bool, error) {
	updateTaskType := &TaskType{
		TaskTypeId:   task_type_id,
		TaskTypeName: task_type_name,
	}

	_, err := db.NewUpdate().Model(updateTaskType).Column("task_type_name").WherePK().Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("error updating task type: %w", err)
	}

	return true, nil
}
