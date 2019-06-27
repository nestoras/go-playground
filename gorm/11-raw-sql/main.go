package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"strconv"
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

type Result struct {
	Name  string
	Email string
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

	for i := 0; i < 10; i++ {
		db.Create(&User{Name: "Nestoras" + strconv.Itoa(i), Email: "nestorasst@gmail.com" + strconv.Itoa(i)})
	}

	var result []Result
	db.Raw("select name, email from users where deleted_at IS NULL").Scan(&result)
	fmt.Println(result)

}
