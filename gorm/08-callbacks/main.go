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
	db.Callback().Create().Register("custom_callback", customCallback)

	var u User

	u.Email = "nestorasst@gmail.com"
	u.Name = "Nestoras Stefanou"

	db.Save(&u)
}

func (u *User) BeforeSave() (err error) {
	fmt.Print("before save hook")
	return nil
}

func (u *User) AfterCreate(scope *gorm.Scope) (err error) {
	fmt.Print("after save hook")
	return nil
}

func customCallback(scope *gorm.Scope) {
	if scope.HasColumn("Email") {
		fmt.Println("Has Column Email")
	} else if scope.HasColumn("lalalalaal") {
		fmt.Println("Don't print this")
	}
}
