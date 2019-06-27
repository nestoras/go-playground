package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/qor/transition"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "golang_course"
)

type Photo struct {
	ID uint
	transition.Transition
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
	db.DropTableIfExists(&Photo{})
	db.AutoMigrate(&Photo{})

	var PhotoStateMachine = transition.New(&Photo{})

	// Define initial state
	PhotoStateMachine.Initial("draft")
	// Define more States
	PhotoStateMachine.State("pending")
	PhotoStateMachine.State("active")
	PhotoStateMachine.State("deleted")

	// Define another State and what to do when entering and exiting that state.
	PhotoStateMachine.State("paid").Enter(func(photo interface{}, tx *gorm.DB) error {
		// To get order object use 'order.(*Order)'
		// business logic here
		fmt.Println("Hello")
		return nil
	}).Exit(func(order interface{}, tx *gorm.DB) error {
		// business logic here
		return nil
	})

	// Define an Event
	PhotoStateMachine.Event("checkout").To("checkout").From("draft")

	// Define another event and what to do before and after performing the transition.
	PhotoStateMachine.Event("pending").To("pending").From("draft").Before(func(order interface{}, tx *gorm.DB) error {
		// business logic here
		fmt.Println("!!!!!!!")
		return nil
	}).After(func(order interface{}, tx *gorm.DB) error {
		// business logic here
		return nil
	})

	// Different state transitions for one event
	cancellEvent := PhotoStateMachine.Event("cancel")
	cancellEvent.To("cancelled").From("draft", "checkout")
	cancellEvent.To("paid_cancelled").From("paid").After(func(order interface{}, tx *gorm.DB) error {
		// Refund
		return nil
	})

	var p Photo

	p.ID = 1

	db.Save(&p)

	PhotoStateMachine.Trigger("pending", &p, db)
	PhotoStateMachine.Trigger("pending", &p, db)

}
