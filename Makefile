# Makefile for task tracker

# Variables
GO_CMD = go
GO_BUILD_CMD = $(GO_CMD) build
GO_TEST_CMD = $(GO_CMD) test
GO_RUN_CMD = $(GO_CMD) run

# Targets
build:
	$(GO_BUILD_CMD) -o bin/task_tracker cmd/main.go

test:
	$(GO_TEST_CMD) ./... -v --cover