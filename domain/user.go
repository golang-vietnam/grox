package domain

type User struct {
	Common
	Username string `json:"username" gorethink:"username"`
	Password string `json:"password" gorethink:"password`
	Name     string `json:"name" gorethink:"name"`
	Email    string `json:"email" gorethink:"email"`
}
