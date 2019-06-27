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

	var c Customer
	db.First(&c)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(c)

	var customer Customer

	customer.Email = "nestorasst@gmail.com"
	db.Where(customer).First(&customer)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(customer)

	//Read all customers
	var customers []Customer
	db.Find(&customers)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println("Total Customers", len(customers))
	fmt.Println(customers)

}
