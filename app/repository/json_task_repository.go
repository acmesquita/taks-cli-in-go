package repository

import (
	"encoding/json"
	"os"

	"github.com/acmesquita/task_tracker/app/model"
)

type JSONTaskRepository struct {
	FilePath string
}

func (r *JSONTaskRepository) loadTasks() []*model.Task {
	if _, err := os.Stat(r.FilePath); os.IsNotExist(err) {
		os.Create(r.FilePath)
		return []*model.Task{}
	}
	jsonData, err := os.ReadFile(r.FilePath)
	if err != nil {
		return []*model.Task{}
	}
	var tasks []*model.Task
	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		return []*model.Task{}
	}
	return tasks
}

func (r *JSONTaskRepository) saveTasks(tasks []*model.Task) {
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		return
	}
	err = os.WriteFile(r.FilePath, jsonData, 0644)
	if err != nil {
		return
	}
}

func (r *JSONTaskRepository) AddTask(task *model.Task) {
	tasks := r.loadTasks()
	tasks = append(tasks, task)
	r.saveTasks(tasks)
}

func (r *JSONTaskRepository) UpdateTask(task *model.Task) {
	tasks := r.loadTasks()
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
		}
	}
	r.saveTasks(tasks)
}

func (r *JSONTaskRepository) DeleteTask(task *model.Task) {
	tasks := r.loadTasks()
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	r.saveTasks(tasks)
}

func (r *JSONTaskRepository) ListTasks() []*model.Task {
	return r.loadTasks()
}

func (r *JSONTaskRepository) GetTask(id string) *model.Task {
	tasks := r.loadTasks()
	for _, t := range tasks {
		if t.ID == id {
			return t
		}
	}
	return nil
}
