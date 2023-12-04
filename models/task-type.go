package models

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

type TaskType struct {
	bun.BaseModel `bun:"table:task_type"`
	TaskTypeid    int    `bun:"pk, autoincrement"`
	TaskTypeName  string `bun:"task_type_name"`
}

func (TaskType) Create_Task_Type(ctx context.Context, task_type_name string, db *bun.DB) (int, error) {
	taskType := &TaskType{
		TaskTypeName: task_type_name,
	}

	_, err := db.NewInsert().Model(taskType).Exec(ctx)
	if err != nil {
		return -1, fmt.Errorf("error inserting task type: %w", err)
	}

	return taskType.TaskTypeid, nil
}

func (TaskType) Update_task_Type(ctx context.Context, id int, task_type_name string, db *bun.DB) (bool, error) {
	updateTaskType := &TaskType{
		TaskTypeid:   id,
		TaskTypeName: task_type_name,
	}

	_, err := db.NewUpdate().Model(updateTaskType).Column("task_type_name").WherePK().Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("error updating task type: %w", err)
	}

	return true, nil
}
