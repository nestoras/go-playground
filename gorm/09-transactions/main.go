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

//
//type User struct {
//	gorm.Model
//	Name   string
//	Email  string
//}

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

	var users []User

	tx := db.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}

	if err := tx.Create(&User{Name: "Nestoras", Email: "nestorasst@gmail.com"}).Error; err != nil {
		tx.Rollback()
	}

	if err := tx.Create(&User{Name: "George", Email: "nestorasst@gmail.com"}).Error; err != nil {
		tx.Rollback()
	}

	tx.Commit()
	if tx.Error != nil {
		tx.Rollback()
	}

	db.Find(&users)

	fmt.Println(users)

}
