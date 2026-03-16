package model

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	task := NewTask("Buy groceries")
	if task.ID == "" {
		t.Errorf("Task ID is empty")
	}
}

func TestValidateTask(t *testing.T) {
	task := NewTask("")
	if task.Validate() {
		t.Errorf("Task is valid")
	}
}

func TestToJSON(t *testing.T) {
	task := NewTask("Buy groceries")
	json := task.ToJSON()
	if json == "" {
		t.Errorf("Task JSON is empty")
	}
}

func TestFromJSON(t *testing.T) {
	json := `{"id": "1", "description": "Buy groceries", "status": "todo", "createdAt": "2021-01-01T00:00:00Z", "updatedAt": "2021-01-01T00:00:00Z"}`
	task := FromJSON(json)

	if task.ID != "1" {
		t.Errorf("Task ID is not 1")
	}
	if task.Description != "Buy groceries" {
		t.Errorf("Task Description is not Buy groceries")
	}
	if task.Status != "todo" {
		t.Errorf("Task Status is not todo")
	}
	if task.CreatedAt != "2021-01-01T00:00:00Z" {
		t.Errorf("Task CreatedAt is not 2021-01-01T00:00:00Z")
	}
	if task.UpdatedAt != "2021-01-01T00:00:00Z" {
		t.Errorf("Task UpdatedAt is not 2021-01-01T00:00:00Z")
	}
}

func TestUpdateTask(t *testing.T) {
	task := NewTask("Buy groceries")
	task.Update("Buy groceries and cook dinner")
	if task.Description != "Buy groceries and cook dinner" {
		t.Errorf("Task Description is not Buy groceries and cook dinner")
	}
}

func TestMarkAsDoneTask(t *testing.T) {
	task := NewTask("Buy groceries")
	task.MarkAsDone()
}

func TestMarkAsInProgressTask(t *testing.T) {
	task := NewTask("Buy groceries")
	task.MarkAsInProgress()
	if task.Status != "in_progress" {
		t.Errorf("Task Status is not in_progress")
	}
}
