package model

type Register struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Login struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
