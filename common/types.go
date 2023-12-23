package common

import "github.com/uptrace/bun"

var DB *bun.DB

type TaskType struct {
	TaskTypeId   int    `json:"task_type_id"`
	TaskTypeName string `json:"task_type_name"`
}
