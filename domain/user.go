package domain

type User struct {
	Common
	Username string `json:"username" gorethink:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" gorethink:"password validate:"required,min=8,max=50"`
	Name     string `json:"name" gorethink:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" gorethink:"email" validate:"required,min=3,max=50,email"`
}
