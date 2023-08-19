package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/katsuomi/gin-like-twitter-api/models"
)

type UserRepository struct{}

type User models.User

func (_UserRepository) GetAll() (User, error) {
	return nil, nil
}
func (_UserRepository) GetById(id int) (User, error) {
	return nil, nil
}
func (_UserRepository) UpdateById(id int, c *gin.Context) {
	return nil, nil
}
func (_UserRepository) DeleteById(id int) error {
	return nil, nil
}
