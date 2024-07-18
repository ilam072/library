package model

type User struct {
	ID           int
	Username     string
	PasswordHash string
	Email        string
}

type UpdateUserDTO struct {
	Username string `json:"username" binding:"required" validate:"max=32"`
	Password string `json:"password" binding:"required" validate:"min=6"`
	Email    string `json:"email" binding:"required" validate:"email"`
}

type RegisterUserDTO struct {
	Username string `json:"username" binding:"required" validate:"max=32"`
	Password string `json:"password" binding:"required" validate:"min=6"`
	Email    string `json:"email" binding:"required" validate:"email"`
}

type SignInUser struct {
	Username string `json:"username" binding:"required" validate:"max=32"`
	Password string `json:"password" binding:"required" validate:"min=6"`
}

type GetUserDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
