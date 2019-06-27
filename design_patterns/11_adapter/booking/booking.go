package booking

import (
	"fmt"
)

type Booking struct {
	Token    string
	Username string
	Password string
}

func (b *Booking) SyncRoom() error {
	fmt.Println("Communicate with Booking to sync rooms")
	return nil
}
