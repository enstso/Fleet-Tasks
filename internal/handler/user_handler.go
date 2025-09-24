package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	userEntity "github.com/enstso/Fleet-Tasks.git/internal/domain/user"
	"github.com/enstso/Fleet-Tasks.git/internal/service"
	"github.com/enstso/Fleet-Tasks.git/internal/utils"
)

// ---------------------- HANDLER: GET ALL USERS ----------------------
func GetUsersHandler(resWriter http.ResponseWriter, req *http.Request) {
	// Only allow GET requests
	if req.Method != http.MethodGet {
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	// Call service layer to fetch all users
	users, err := service.GetUsers()
	if err != nil {
		// If service fails, stop (better to return a 500 error)
		return
	}

	// Set response type to JSON
	resWriter.Header().Set(utils.HeaderContentTypeValue, utils.HeaderContentTypeValue)

	// Encode user list as JSON and write response
	err = json.NewEncoder(resWriter).Encode(users)
	if err != nil {
		// If encoding fails, stop silently (better to return 500 error)
		return
	}
}

// ---------------------- HANDLER: GET USER BY ID ----------------------
func GetUserByIdHandler(resWriter http.ResponseWriter, req *http.Request) {
	// Only allow GET requests
	if req.Method != http.MethodGet {
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	// Split URL path: /user/{id}
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 3 {
		// If no ID is provided, return 400 Bad Request
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusBadRequest)
		return
	}

	// Extract ID from the URL
	id := parts[2]

	// Call service layer to get a user by ID
	user, err := service.GetUserById(id)
	if err != nil {
		// If service fails, return 500 Internal Server Error
		http.Error(resWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response type to JSON
	resWriter.Header().Set("Content-Type", "application/json")

	// Encode single user as JSON and send response
	err = json.NewEncoder(resWriter).Encode(user)
	if err != nil {
		return
	}
}

// ---------------------- HANDLER: CREATE USER ----------------------
func CreateUserHandler(resWriter http.ResponseWriter, req *http.Request) {
	// Only allow POST requests
	if req.Method != http.MethodPost {
		http.Error(resWriter, utils.ErrNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	// Decode request body into User DTO
	var user userEntity.Dto
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		// If request body is invalid JSON, stop (better: return 400 error)
		return
	}

	// Call service layer to create a new user
	service.CreateUser(user)

	// Should return 201 Created + the created user in JSON
	resWriter.Header().Set(utils.HeaderContentTypeValue, utils.HeaderContentTypeValue)
	resWriter.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(resWriter).Encode(user)
}
