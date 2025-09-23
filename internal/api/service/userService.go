package service

import (
	"slices"

	"github.com/enstso/Fleet-Tasks.git/internal/api"
	"github.com/enstso/Fleet-Tasks.git/internal/api/domain/user"
)

var users []user.User

func getUsers() []user.Dto {
	var usersDtoJson []user.Dto
	if len(users) == 0 {
		return usersDtoJson
	}
	for _, v := range users {
		usersDtoJson = append(usersDtoJson, user.ToDto(v))
	}
	return usersDtoJson
}

func getUserById(id string) (user.Dto, error) {
	var userFound user.User
	var idExist = slices.IndexFunc(
		users, func(user user.User) bool {
			if user.ID == id {
				userFound = user
				return true
			}
			return false
		})
	if idExist == -1 {
		return user.Dto{}, api.ErrNotFound
	}
	return user.ToDto(userFound), nil
}

func createUser(userDto user.Dto) {
	var userObj user.User
	userObj = user.FromDto(userDto)
	users = append(users, userObj)
}
