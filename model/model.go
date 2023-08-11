package main

type User struct {
	ID       uint   `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Post struct {
	ID    uint   `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Title string `json:"title" binding:"required"`
	User  User   `json:"-" binding:"required"`
}
