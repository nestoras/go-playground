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

type Customer struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Photos []Photo
}

type Photo struct {
	gorm.Model
	CustomerID uint
	Filename   string
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
	db.AutoMigrate(&Customer{}, &Photo{})

	var customer Customer
	db.Preload("Photos").First(&customer)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println("Email:", customer.Email)
	fmt.Println("Total photos:", len(customer.Photos))
	fmt.Println("Photos:", customer.Photos)
}
