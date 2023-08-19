package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Todo struct {
	gorm.Model
	Name   string `json:"name" binding:"required"`
	Title  string `json:"title" binding:"required"`
	UserId uint   `json:"-" binding:"required"`
}

type SignUpInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"name" binding:"required"`
	Password string `json:"name" binding:"required"`
}

func Encrypt(char string) string {
	encryptText := fmt.Sprintf("%x", sha256.Sum256([]byte(char)))
	return encryptText
}

func (user *User) Create(db *gorm.DB) (User, error) {
	user := User{
			Name: user.Name,
			Email: user.Email,
			Password: Encrypt(user.Password)
	}
	result := db.Create(&user)

	return user, result.Error
}
