package repository

import "github.com/acmesquita/task_tracker/app/model"

type TaskRepositoryImpl struct {
	Tasks []*model.Task
}

func (r *TaskRepositoryImpl) AddTask(task *model.Task) {
	r.Tasks = append(r.Tasks, task)
}

func (r *TaskRepositoryImpl) UpdateTask(task *model.Task) {
	for i, t := range r.Tasks {
		if t.ID == task.ID {
			r.Tasks[i] = task
		}
	}
}

func (r *TaskRepositoryImpl) DeleteTask(task *model.Task) {
	for i, t := range r.Tasks {
		if t.ID == task.ID {
			r.Tasks = append(r.Tasks[:i], r.Tasks[i+1:]...)
		}
	}
}

func (r *TaskRepositoryImpl) ListTasks() []*model.Task {
	return r.Tasks
}

func (r *TaskRepositoryImpl) GetTask(id string) *model.Task {
	for _, t := range r.Tasks {
		if t.ID == id {
			return t
		}
	}
	return nil
}
