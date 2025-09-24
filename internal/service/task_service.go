package service

import (
	"slices"
	"strconv"

	taskEntity "github.com/enstso/Fleet-Tasks.git/internal/domain/task"
	"github.com/enstso/Fleet-Tasks.git/internal/utils"
)

// In-memory slice to store tasks (non-persistent, lost when program stops)
var tasks []taskEntity.Task

// ---------------------- GET ALL TASKS ----------------------
// GetTasks converts all stored Task entities into DTOs and returns them.
// Returns an empty slice if no tasks exist.
func GetTasks() ([]taskEntity.Dto, error) {
	var tasksDto []taskEntity.Dto

	// If there are no tasks, return an empty slice (no error)
	if len(tasks) == 0 {
		return tasksDto, nil
	}

	// Convert each Task entity into a DTO
	for _, v := range tasks {
		taskDto, err := taskEntity.ToDto(v)
		if err != nil {
			// If conversion fails, return an error
			return []taskEntity.Dto{}, err
		}
		tasksDto = append(tasksDto, taskDto)
	}

	return tasksDto, nil
}

// ---------------------- GET TASK BY ID ----------------------
// GetTaskById finds a task in the slice by its ID and returns it as a DTO.
// If not found, returns a NotFound error.
func GetTaskById(id string) (taskEntity.Dto, error) {
	var taskFound taskEntity.Task

	// Search tasks slice for a matching ID
	idExist := slices.IndexFunc(tasks, func(task taskEntity.Task) bool {
		if task.ID == id {
			taskFound = task // Save the found task
			return true
		}
		return false
	})

	// If ID not found, return "not found" error
	if idExist == -1 {
		return taskEntity.Dto{}, utils.ErrNotFound
	}

	// Convert the found task to a DTO (ignoring conversion error here)
	taskJson, _ := taskEntity.ToDto(taskFound)
	return taskJson, nil
}

// ---------------------- CREATE TASK ----------------------
// CreateTask takes a DTO, assigns it a new ID, and converts it to an entity.
// NOTE: The current implementation does NOT append the task to the slice,
// so the task is never actually saved in memory.
func CreateTask(dto taskEntity.Dto) {

	// Get last task ID (int)
	lastIdTaskFound, _ := lastIdTask()

	// Assign new ID (currently reuses last ID instead of last+1)
	dto.ID = strconv.Itoa(lastIdTaskFound)

	// Convert DTO into Task entity (ignore result and error)
	task, err := taskEntity.FromDto(dto)
	if err != nil {
		return
	}
	tasks = append(tasks, task)
}

// ---------------------- DELETE TASK ----------------------
// DeleteTask removes a task with the given ID from the slice.
func DeleteTask(id string) {
	// Find index of task to delete
	positionExist := slices.IndexFunc(tasks, func(task taskEntity.Task) bool {
		return task.ID == id
	})

	// Delete the slice element at found position
	// NOTE: slices.Delete returns a new slice,
	// but the result is ignored here — so nothing is actually deleted.
	slices.Delete(tasks, positionExist, positionExist+1)
}

// ---------------------- LAST TASK ID ----------------------
// lastIdTask returns the numeric ID of the last task in the slice.
// If the slice is empty or ID cannot be converted to int, returns -1 and error.
func lastIdTask() (int, error) {

	if len(tasks) == 0 {
		return 1, nil
	}

	// Get the last task in the slice (will panic if slice is empty)
	lastTask := tasks[len(tasks)-1]

	// Try to parse ID as int
	if id, err := strconv.Atoi(lastTask.ID); err == nil {
		return id, nil
	}
	return -1, utils.ErrNotExist
}
