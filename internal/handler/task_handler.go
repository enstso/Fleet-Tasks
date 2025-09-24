package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	taskEntity "github.com/enstso/Fleet-Tasks.git/internal/domain/task"
	"github.com/enstso/Fleet-Tasks.git/internal/service"
	"github.com/enstso/Fleet-Tasks.git/internal/utils"
)

// ---------------------- HANDLER: GET ALL TASKS ----------------------
func GetTasksHandler(resWriter http.ResponseWriter, req *http.Request) {
	// Only allow GET requests
	if req.Method != http.MethodGet {
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	// Call service layer to get all tasks
	tasks, err := service.GetTasks()
	if err != nil {
		// If service fails, just return (better to log + send error)
		println(err)
		return
	}

	// Set response header to JSON
	resWriter.Header().Set(utils.HeaderContentTypeValue, utils.HeaderContentTypeValue)

	// Encode tasks as JSON and write to response
	err = json.NewEncoder(resWriter).Encode(tasks)
	if err != nil {
		// If encoding fails, stop silently (better to handle explicitly)
		return
	}
}

// ---------------------- HANDLER: GET TASK BY ID ----------------------
func GetTasksByIdHandler(resWriter http.ResponseWriter, req *http.Request) {
	// Only allow GET requests
	if req.Method != http.MethodGet {
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	// Split URL path: /task/{id}
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 3 {
		// Not enough parts, meaning ID is missing
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusBadRequest)
		return
	}

	// Extract ID from path
	id := parts[2]

	// Call service layer to get a single task by ID
	task, err := service.GetTaskById(id)
	if err != nil {
		// If something goes wrong, return 500
		http.Error(resWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header to JSON
	resWriter.Header().Set("Content-Type", "application/json")

	// Encode single task as JSON
	err = json.NewEncoder(resWriter).Encode(task)
	if err != nil {
		return
	}
}

// DeleteTaskHandler handles HTTP DELETE requests to remove a task by its ID.
// Expected path: /task/{id}
// - Returns 405 Method Not Allowed if the method is not DELETE.
// - Returns 400 Bad Request if no ID is provided in the URL.
// - Calls the service layer to delete the task with the given ID.
// - Currently always responds with 200 OK and null JSON (improvement: 204 No Content or 404 if not found).

func DeleteTaskHandler(resWriter http.ResponseWriter, req *http.Request) {
	// Only allow DELETE requests
	if req.Method != http.MethodDelete {
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	// Split URL path: /task/{id}
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 3 {
		// Not enough parts, meaning ID is missing
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusBadRequest)
		return
	}

	// Extract ID from path
	id := parts[2]

	// Call service layer to delete a task with this ID
	service.DeleteTask(id)

	// Set response header to JSON
	resWriter.Header().Set(utils.HeaderContentTypeValue, utils.HeaderJSONValue)
	resWriter.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(resWriter).Encode("Deleted")
}

// CreateTaskHandler ---------------------- HANDLER: CREATE TASK ----------------------
func CreateTaskHandler(resWriter http.ResponseWriter, req *http.Request) {
	// Only allow POST requests
	if req.Method != http.MethodPost {
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	// Decode request body into Task DTO
	var task taskEntity.Dto
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		// If decoding fails (bad JSON), just return (better: send 400 error)
		println(err.Error())
		return
	}

	// Call service layer to create the new task
	service.CreateTask(task)

	// Should return 201 Created + the created task in JSON
	resWriter.Header().Set(utils.HeaderContentTypeValue, utils.HeaderContentTypeValue)
	resWriter.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(resWriter).Encode(task)
}
