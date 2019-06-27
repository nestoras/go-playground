package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
)

type StatusType string

const (
	Active   StatusType = "active"
	Inactive StatusType = "inactive"
	Pending  StatusType = "pending"
	Delete   StatusType = "delete"
)

func (u *StatusType) Scan(value interface{}) error { *u = StatusType(value.([]byte)); return nil }
func (u StatusType) Value() (driver.Value, error)  { return string(u), nil }

type Photo struct {
	Filename string     `gorm:"not null;type:varchar(100);unique_index"`
	Status   StatusType `gorm:"not null;type:ENUM('active', 'inactive', 'pending', 'delete')" json:"-"`
}

func main() {
	db, err := gorm.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/golang_course?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.DropTableIfExists(&Photo{})
	db.AutoMigrate(&Photo{})

	err = db.Create(Photo{Filename: "my-photo.jpg", Status: "pending"}).Error

	if err != nil {
		panic(err)
	}

	var photos []Photo
	err = db.Find(&photos).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(photos)
}
