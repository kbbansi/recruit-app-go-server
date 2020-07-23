package models

type Auth struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthSuccess struct {
	Status int `json:"status"`
	Message string `json:"message"`
	SessionData *User `json:"session_data"`
}