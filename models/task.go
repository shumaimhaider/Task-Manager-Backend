package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/shumaimhaider/task_manager_api/common"
	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"todo_task"`
	ID            int      `bun:"task_id,pk,autoincrement"`
	Details       string   `bun:"details"`
	IsCompleted   bool     `bun:"is_completed"`
	TaskTypeID    int      `bun:"task_type_id"`
	TaskType      TaskType `bun:"rel:belongs-to,join:task_id=task_type_id"`
}

func CreateTask(ctx context.Context, details string, iscompleted bool, taskTypeID int) (int, error) {
	taskType := &Task{Details: details, IsCompleted: iscompleted, TaskTypeID: taskTypeID}
	_, err := common.DB.NewInsert().Model(taskType).Exec(ctx)
	if err != nil {
		return -1, fmt.Errorf("error creating task %v", err)
	}
	return taskType.ID, nil
}

func UpdateTaskCompleted(ctx context.Context, id int, iscompleted bool) (bool, error) {
	task := &Task{ID: id, IsCompleted: iscompleted}
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

func GetAllToDosTask(ctx context.Context, db *bun.DB) ([]common.Task, error) {
	var toDoTasks []Task

	err := db.NewSelect().Model(&toDoTasks).Relation("ToDoType").OrderExpr("todo_task_id ASC").Scan(ctx)

	// err := db.NewSelect().Model(&toDoTasks).Relation("ToDoType").Where("is_completed = ?", true).OrderExpr("todo_task_id ASC").Scan(context.Background())
	// err := db.NewSelect().Model(&toDoTasks).
	// 	Relation("ToDoType").
	// 	GroupExpr("todo_type_id").
	// 	Having("Min(is_completed=true)>1").
	// 	Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error get all tasks %v", err)
	}
	return convertDBTodoToControllerTodo(toDoTasks), nil
}

// converter
func convertDBTodoToControllerTodo(dbTodos []Task) []common.Task {
	result := make([]common.Task, len(dbTodos))
	for i, dbTodo := range dbTodos {
		controllerTodo := common.Task{
			ID:          dbTodo.ID,
			Details:     dbTodo.Details,
			IsCompleted: dbTodo.IsCompleted,
			TaskTypeID:  dbTodo.TaskTypeID,
			TaskType: common.TaskType{
				TaskTypeId:   dbTodo.TaskType.TaskTypeId,
				TaskTypeName: dbTodo.TaskType.TaskTypeName,
			},
		}
		result[i] = controllerTodo
	}
	return result
}

func DeleteTask(ctx context.Context, todo_task_id int) (bool, error) {
	task := &Task{ID: todo_task_id}
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
