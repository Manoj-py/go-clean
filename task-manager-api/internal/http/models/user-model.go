package models

type UserRequestModel struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	Role     string `json:"role" validate:"required, oneof=admin,user"`
}
