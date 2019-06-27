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
	db.AutoMigrate(&Customer{})

	// Create
	db.Create(&Customer{Name: "Test", Email: "test@test.com"})

	// Read
	var customer Customer
	db.First(&customer, "email = ?", "test@test.com") //
	fmt.Print(customer)

	// Update customer's name
	db.Model(&customer).Update("Name", "Test2")
	fmt.Print(customer)
	// Delete - delete product
	//db.Delete(&customer)
	//fmt.Print(&customer)

}
