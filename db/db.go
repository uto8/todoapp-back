package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/katsuomi/gin-like-twitter-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	dsn := "root:password@tcp(mysql_db)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigration()
	user := models.User{
		ID:    1,
		Name:  "aoki",
		Posts: []models.Post{{ID: 1, Content: "tweet1"}, {ID: 2, Content: "tweet2"}},
	}
	db.Create(&user)
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}
