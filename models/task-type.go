package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shumaimhaider/task_manager_api/common"
	"github.com/uptrace/bun"
)

type TaskType struct {
	bun.BaseModel `bun:"table:task_type"`
	TaskTypeId    int    `bun:"task_type_id,pk,autoincrement"`
	TaskTypeName  string `bun:"task_type_name,notempty"`
}

func Create_Task_Type(ctx *gin.Context, task_type_name string) (int, error) {
	taskType := &TaskType{
		TaskTypeName: task_type_name,
	}

	_, err := common.DB.NewInsert().Model(taskType).Exec(ctx)
	if err != nil {
		return -1, fmt.Errorf("error inserting task type: %w", err)
	}

	return taskType.TaskTypeId, nil
}

func Update_task_Type(ctx *gin.Context, task_type_id int, task_type_name string) (bool, error) {
	updateTaskType := &TaskType{
		TaskTypeId:   task_type_id,
		TaskTypeName: task_type_name,
	}

	_, err := common.DB.NewUpdate().Model(updateTaskType).Column("task_type_name").WherePK().Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("error updating task type: %w", err)
	}

	return true, nil
}

func Delete_ToDo_Type(ctx context.Context, taskTypeName string) (bool, error) {
	TaskType := &TaskType{TaskTypeName: taskTypeName}
	res, err := common.DB.NewDelete().Model(TaskType).WherePK().Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("error deleting task type type: %w", err)
	}
	rowsAffected, err := res.RowsAffected()

	switch {
	case err != nil:
		return false, fmt.Errorf("error getting rows affetced: %w", err)
	case rowsAffected == 0:
		return false, errors.New("task type not found")
	}

	return true, nil
}

func GetAllTaskType(ctx context.Context) ([]common.TaskType, error) {
	var TaskType []TaskType
	err := common.DB.NewSelect().Model(&TaskType).OrderExpr("task_type_id ASC").Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting all task types %v", err)
	}
	return converterTaskType(TaskType), nil
}

func converterTaskType(dbTaskType []TaskType) []common.TaskType {
	result := make([]common.TaskType, len(dbTaskType))
	for i, dbTaskType := range dbTaskType {
		taskType := common.TaskType{
			TaskTypeId:   dbTaskType.TaskTypeId,
			TaskTypeName: dbTaskType.TaskTypeName,
		}
		result[i] = taskType
	}

	return result
}
