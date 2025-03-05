package dto

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	ID int `json:"id"`
}
