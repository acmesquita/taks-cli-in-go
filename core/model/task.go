package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

const (
	TODO        = "todo"
	IN_PROGRESS = "in_progress"
	DONE        = "done"
)

func NewTask(description string) *Task {
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	return &Task{
		ID:          id,
		Description: description,
		Status:      TODO,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
}

func (t *Task) Validate() bool {
	return t.Description != ""
}

func (t *Task) ToJSON() string {
	json, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(json)
}

func FromJSON(jsonData string) *Task {
	var task Task
	err := json.Unmarshal([]byte(jsonData), &task)
	if err != nil {
		return nil
	}
	return &task
}

func (t *Task) Update(description string) {
	t.Description = description
	t.UpdatedAt = time.Now().Format(time.RFC3339)
}

func (t *Task) MarkAsDone() {
	t.Status = DONE
	t.UpdatedAt = time.Now().Format(time.RFC3339)
}

func (t *Task) MarkAsInProgress() {
	t.Status = IN_PROGRESS
	t.UpdatedAt = time.Now().Format(time.RFC3339)
}
