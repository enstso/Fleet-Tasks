package service

import (
	"slices"
	"strconv"

	"github.com/enstso/Fleet-Tasks.git/internal/api"
	"github.com/enstso/Fleet-Tasks.git/internal/api/domain/task"
)

var tasks []task.Task

func getTasks() ([]task.Dto, error) {
	var tasksDto []task.Dto
	for _, v := range tasks {
		taskDto, err := task.ToDto(v)
		if err != nil {
			return []task.Dto{}, err
		}
		tasksDto = append(tasksDto, taskDto)
	}
	return tasksDto, nil
}

func getTaskById(id string) (task.Dto, error) {
	var taskFound task.Task
	var idExist = slices.IndexFunc(tasks, func(task task.Task) bool {
		if task.ID == id {
			taskFound = task
			return true
		}
		return false
	})

	if idExist == -1 {
		return task.Dto{}, api.ErrNotFound
	}
	var taskJson, _ = task.ToDto(taskFound)

	return taskJson, nil

}

func createTask(dto task.Dto) {
	var lastIdTaskFound, _ = lastIdTask()

	dto.ID = strconv.Itoa(lastIdTaskFound)

	_, err := task.FromDto(dto)
	if err != nil {
		return
	}
}

func deleteTask(id string) {
	var positionExist = slices.IndexFunc(tasks, func(task task.Task) bool {
		if task.ID == id {
			return true
		}
		return false
	})
	slices.Delete(tasks, positionExist, positionExist+1)
}

func lastIdTask() (int, error) {
	var lastTask = tasks[len(tasks)-1]
	if id, err := strconv.Atoi(lastTask.ID); err == nil {
		return id, nil
	}
	return -1, api.ErrNotExist
}
