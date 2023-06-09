package models

type SignupJson struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
