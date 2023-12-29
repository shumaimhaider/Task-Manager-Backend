package common

import "github.com/uptrace/bun"

var DB *bun.DB

type TaskType struct {
	TaskTypeId   int    `json:"task_type_id"`
	TaskTypeName string `json:"task_type_name"`
}

type Task struct {
	ID          int    `json:"task_id"`
	Details     string `json:"details"`
	IsCompleted bool   `json:"is_completed"`
	TaskTypeID  int    `json:"task_type_id"`
	TaskType    TaskType
}
