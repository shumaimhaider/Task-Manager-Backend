package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/shumaimhaider/task_manager_api/common"
	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"task"`
	ID            int      `bun:"task_id,pk,autoincrement"`
	Details       string   `bun:"details"`
	IsCompleted   bool     `bun:"is_completed"`
	TaskTypeID    int      `bun:"task_type_id"`
	TaskType      TaskType `bun:"rel:belongs-to,join:task_id=task_type_id"`
}

func CreateTask(ctx context.Context, req common.TaskRequest) (int, error) {
	taskType := &Task{Details: req.Details, IsCompleted: req.IsCompleted, TaskTypeID: req.TaskTypeID}
	_, err := common.DB.NewInsert().Model(taskType).Exec(ctx)
	if err != nil {
		return -1, fmt.Errorf("error creating task %v", err)
	}
	return taskType.ID, nil
}

func UpdateTaskCompleted(ctx context.Context, req common.TaskRequest) (bool, error) {
	task := &Task{ID: req.ID, IsCompleted: req.IsCompleted}
	res, err := common.DB.NewUpdate().Model(task).Set("is_completed = ?", task.IsCompleted).Where("task_id = ?", task.ID).Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("error updating task: %w", err)
	}
	rowsAffected, err := res.RowsAffected()
	switch {
	case err != nil:
		return false, fmt.Errorf("error getting rows affetced: %w", err)
	case rowsAffected == 0:
		return false, errors.New("task not found")
	}

	return true, nil
}

func GetAllTask(ctx context.Context) ([]common.Task, error) {
	var task []Task

	err := common.DB.NewSelect().Model(&task).Relation("TaskType").OrderExpr("task_id ASC").Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error get all tasks %v", err)
	}
	return converterTask(task), nil
}

func converterTask(dbTask []Task) []common.Task {
	result := make([]common.Task, len(dbTask))
	for i, dbTask := range dbTask {
		controllerTodo := common.Task{
			ID:          dbTask.ID,
			Details:     dbTask.Details,
			IsCompleted: dbTask.IsCompleted,
			TaskTypeID:  dbTask.TaskTypeID,
			TaskType: common.TaskType{
				TaskTypeId:   dbTask.TaskType.TaskTypeId,
				TaskTypeName: dbTask.TaskType.TaskTypeName,
			},
		}
		result[i] = controllerTodo
	}
	return result
}

func DeleteTask(ctx context.Context, task_id int) (bool, error) {
	task := &Task{ID: task_id}
	res, err := common.DB.NewDelete().Model(task).WherePK().Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("error deleting task %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	switch {
	case err != nil:
		return false, fmt.Errorf("error getting rows affetced: %w", err)
	case rowsAffected == 0:
		return false, errors.New("task not found")
	}

	return true, nil
}
