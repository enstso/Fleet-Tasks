package user

type Dto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToDto(user User) Dto {
	return Dto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func FromDto(userDto Dto) User {
	return User{
		ID:    userDto.ID,
		Name:  userDto.Name,
		Email: userDto.Email,
	}
}
