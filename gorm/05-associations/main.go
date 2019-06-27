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

type Photo struct {
	gorm.Model
	CustomerID uint
	Filename   string
}

type Customer struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Photos []Photo
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	must(err)

	defer db.Close()
	db.LogMode(true)
	db.DropTableIfExists(&Photo{}, Customer{})
	db.AutoMigrate(&Photo{}, Customer{})

	err = createUser(db, "nestoras", "nestorasst@gmail.com")
	must(err)
	err = createUser(db, "nestoras2", "nestorasst2@gmail.com")
	must(err)
	err = createUser(db, "nestoras3", "nestorasst3@gmail.com")
	must(err)

	var cFirst Customer
	db.First(&cFirst)
	fmt.Println(cFirst)

	var cLast Customer
	db.Last(&cLast)
	fmt.Println(cLast)

	var customer Customer
	db.First(&customer, 2)
	fmt.Println(customer)

	createPhoto(db, cFirst, "random-photo1.jpg")
	createPhoto(db, cFirst, "random-photo111.jpg")
	createPhoto(db, cFirst, "random-photo122.jpg")
	createPhoto(db, cFirst, "random-photo133.jpg")
	createPhoto(db, cLast, "random-photo2.jpg")
	createPhoto(db, customer, "random-photo3.jpg")

}

func createUser(db *gorm.DB, name, email string) error {
	db.Create(&Customer{
		Name:  name,
		Email: email,
	})
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func createPhoto(db *gorm.DB, customer Customer, filename string) error {
	db.Create(&Photo{
		CustomerID: customer.ID,
		Filename:   filename,
	})
	if db.Error != nil {
		return db.Error
	}
	return nil

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
