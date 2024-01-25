package authDomain

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Birthday string `json:"birthday"`
}

type LoginBodyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
