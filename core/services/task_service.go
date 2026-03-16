package services

import (
	"github.com/acmesquita/task_tracker/core/model"
	"github.com/acmesquita/task_tracker/core/repository"
)

type TaskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) *TaskService {
	return &TaskService{taskRepository: taskRepository}
}

func (s *TaskService) AddTask(description string) *model.Task {
	task := model.NewTask(description)
	s.taskRepository.AddTask(task)
	return task
}

func (s *TaskService) UpdateTask(id string, description string) *model.Task {
	task := s.taskRepository.GetTask(id)
	if task == nil {
		return nil
	}
	task.Update(description)
	s.taskRepository.UpdateTask(task)
	return task
}

func (s *TaskService) DeleteTask(id string) *model.Task {
	task := s.taskRepository.GetTask(id)
	if task == nil {
		return nil
	}
	s.taskRepository.DeleteTask(task)
	return task
}

func (s *TaskService) ListTasks() []*model.Task {
	return s.taskRepository.ListTasks()
}

func (s *TaskService) GetTask(id string) *model.Task {
	return s.taskRepository.GetTask(id)
}

func (s *TaskService) MarkTaskAsDone(id string) *model.Task {
	task := s.taskRepository.GetTask(id)
	if task == nil {
		return nil
	}
	task.MarkAsDone()
	s.taskRepository.UpdateTask(task)
	return task
}

func (s *TaskService) MarkTaskAsInProgress(id string) *model.Task {
	task := s.taskRepository.GetTask(id)
	if task == nil {
		return nil
	}
	task.MarkAsInProgress()
	s.taskRepository.UpdateTask(task)
	return task
}

func (s *TaskService) ListTasksByStatus(status string) []*model.Task {
	tasks := s.taskRepository.ListTasks()
	if status == "all" {
		return tasks
	}
	filteredTasks := make([]*model.Task, 0)
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
