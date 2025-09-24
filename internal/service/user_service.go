package service

import (
	"slices"
	"strconv"

	"github.com/enstso/Fleet-Tasks.git/internal/domain/user"
	"github.com/enstso/Fleet-Tasks.git/internal/utils"
)

// In-memory slice to store users (non-persistent, lost when program stops)
var users []user.User

// GetUsers ---------------------- GET ALL USERS ----------------------
// GetUsers converts all stored User entities into DTOs and returns them.
// Returns an empty slice if no users exist.
func GetUsers() ([]user.Dto, error) {
	var usersDtoJson []user.Dto

	// If there are no users, return empty slice (no error)
	if len(users) == 0 {
		return usersDtoJson, nil
	}

	// Convert each User entity to DTO and append to result
	for _, v := range users {
		usersDtoJson = append(usersDtoJson, user.ToDto(v))
	}

	return usersDtoJson, nil
}

// GetUserById ---------------------- GET USER BY ID ----------------------
// GetUserById searches the in-memory slice for a user with the given ID.
// If found, returns it as a DTO. If not found, returns a NotFound error.
func GetUserById(id string) (user.Dto, error) {
	var userFound user.User

	// Find index of user whose ID matches
	idExist := slices.IndexFunc(users, func(user user.User) bool {
		if user.ID == id {
			userFound = user // Save matched user
			return true
		}
		return false
	})

	// If not found, return error
	if idExist == -1 {
		return user.Dto{}, utils.ErrNotFound
	}

	// Return the found user as DTO
	return user.ToDto(userFound), nil
}

// CreateUser ---------------------- CREATE USER ----------------------
// CreateUser takes a DTO, converts it to a User entity, and appends it to the slice.
func CreateUser(userDto user.Dto) {
	// Get last task ID (int)
	lastIdUserFound, _ := lastIdUser()

	userDto.ID = strconv.Itoa(lastIdUserFound)

	if len(users) == 0 {
		userDto.ID = "1"
	}
	var userObj user.User
	// Convert DTO to entity
	userObj = user.FromDto(userDto)

	// Add new user to in-memory slice
	users = append(users, userObj)
}

func lastIdUser() (int, error) {

	if len(users) == 0 {
		return 1, nil
	}

	lastUser := users[len(users)-1]

	if id, err := strconv.Atoi(lastUser.ID); err == nil {
		return id, nil
	}
	return -1, utils.ErrNotExist
}
