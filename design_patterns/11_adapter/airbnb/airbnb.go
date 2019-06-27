package airbnb

import "fmt"

type AirBnb struct {
	APIKey string
}

func (a *AirBnb) Send() error {
	fmt.Println("Communicate with AirBNB to sync rooms")
	return nil
}
