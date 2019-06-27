package main

import (
	"golang-course/design_patterns/11_adapter/airbnb"
	"golang-course/design_patterns/11_adapter/booking"
	"time"
)

type Sync interface {
	SyncRoom(checkIn, checkOut time.Time) error
}

type Checkout struct {
	Reservation Reservation
	SyncType    Sync
}

func (c *Checkout) Checkout() error {
	return c.SyncType.SyncRoom(c.Reservation.CheckIn, c.Reservation.CheckOut)
}

type Reservation struct {
	CheckIn  time.Time
	CheckOut time.Time
	Customer string
}

type BookingAdapter struct {
	Gateway *booking.Booking
}

func (b *BookingAdapter) SyncRoom(checkIn, checkOut time.Time) error {
	b.Gateway.SyncRoom()
	return nil
}

type AirBnbAdapter struct {
	Gateway *airbnb.AirBnb
}

func (p *AirBnbAdapter) SyncRoom(checkIn, checkOut time.Time) error {
	p.Gateway.Send()
	return nil
}

func main() {
	bookingAdapter := BookingAdapter{
		Gateway: &booking.Booking{
			Token:    "booking-token-123",
			Username: "nestoras",
			Password: "123456",
		},
	}

	airbnbAdapter := AirBnbAdapter{
		Gateway: &airbnb.AirBnb{
			APIKey: "airbnb-token-1234",
		},
	}

	reservation := &Reservation{
		CheckIn:  time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
		CheckOut: time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
		Customer: "Nestoras Stefanou",
	}

	checkout := &Checkout{}
	checkout.Reservation = *reservation
	checkout.SyncType = &bookingAdapter
	checkout.Checkout()

	checkout2 := &Checkout{}
	checkout2.Reservation = *reservation
	checkout2.SyncType = &airbnbAdapter
	checkout2.Checkout()
}
