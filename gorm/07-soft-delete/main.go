package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "golang_course"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})

	user := User{Name: "Nestoras", Email: "nestoras-delete@gmail.com"}
	// Create user
	db.Create(&user)

	// Get user
	db.First(&user)

	// Soft delete
	db.Delete(&user)

	// Soft deleted records will be ignored when query them
	db.Where("email =?", "nestoras-delete@gmail.com").Find(&user)

	// Find soft deleted records
	db.Unscoped().Where("email =?", "nestoras-delete@gmail.com").Find(&user)

	// Delete record permanently
	db.Unscoped().Delete(&user)

	// Find again soft deleted records
	db.Unscoped().Where("email =?", "nestoras-delete@gmail.com").Find(&user)
}
