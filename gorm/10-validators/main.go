package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/qor/validations"
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

	//Register callbacks
	validations.RegisterCallbacks(db)

	var errors []error

	errors = db.Create(&User{Name: "N", Email: ""}).GetErrors()
	fmt.Println(errors)

	errors = db.Create(&User{Name: "Nestoras", Email: "nestorasst@gmail.com"}).GetErrors()
	fmt.Println(errors)

	var users []User
	db.Find(&users)
	fmt.Println(users)

}

func (user User) Validate(db *gorm.DB) {
	if user.Email == "" {
		db.AddError(errors.New("Email can't be blank"))
	}

	if len(user.Name) < 2 {
		db.AddError(errors.New("Your name is wrong"))
	}
}
