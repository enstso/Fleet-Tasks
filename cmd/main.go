package main

import (
	"log"
	"net/http"

	"github.com/enstso/Fleet-Tasks.git/internal/handler"
	"github.com/enstso/Fleet-Tasks.git/internal/utils"
)

// ServerApi sets up routes and starts the HTTP server.
func main() {
	// -------- TASKS ROUTES --------

	// GET /tasks → Returns a list of all tasks
	http.HandleFunc("/tasks", handler.GetTasksHandler)

	// POST /task → Creates a new task
	http.HandleFunc("/task", handler.CreateTaskHandler)

	http.HandleFunc("/task/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// GET /task/{id} → Returns a single task by ID (prefix match: /task/123)
			handler.GetTasksByIdHandler(w, r)
			break
		case http.MethodDelete:
			//DELETE /task/{id} -> Delete a task
			handler.DeleteTaskHandler(w, r)
		default:
			http.Error(w, utils.ErrNotAllowed.Error(), http.StatusMethodNotAllowed)
		}
	})

	// -------- USERS ROUTES --------

	// GET /users → Returns a list of all users
	http.HandleFunc("/users", handler.GetUsersHandler)

	// GET /user/{id} → Returns a single user by ID (prefix match: /user/123)
	http.HandleFunc("/user/", handler.GetUserByIdHandler)

	// POST /user → Creates a new user
	http.HandleFunc("/user", handler.CreateUserHandler)

	// -------- SERVER START --------
	// Start the HTTP server on port 8080
	// ":8080" means it will listen on all network interfaces, not just localhost
	// log.Fatal ensures errors are logged if the server fails to start

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
