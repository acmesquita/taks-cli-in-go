package repository

import "github.com/acmesquita/task_tracker/app/model"

type TaskRepository interface {
	AddTask(task *model.Task)
	UpdateTask(task *model.Task)
	DeleteTask(task *model.Task)
	ListTasks() []*model.Task
	GetTask(id string) *model.Task
}
