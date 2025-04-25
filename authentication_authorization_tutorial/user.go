package main

type User struct {
	ID       uint   `json:"_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
