package main

import "fmt"

type Listener struct {
	ID int
}

type ListenerInterface interface {
	execute(m string)
}

func (l *Listener) execute(m string) {
	fmt.Printf("%q message receiver for id %d \n", m, l.ID)
}

type Subject struct {
	listeners []ListenerInterface
}

func (s *Subject) addListener(l ListenerInterface) {
	s.listeners = append(s.listeners, l)
}

func (s *Subject) notify(m string) {
	for _, l := range s.listeners {
		if l != nil {
			l.execute(m)
		}
	}
}

var iterator int

func newListener() *Listener {
	l := Listener{iterator}
	iterator++
	return &l
}

func main() {
	iterator = 0
	s := Subject{listeners: make([]ListenerInterface, 0)}

	l := newListener()
	s.addListener(l)

	for i := 0; i < 5; i++ {
		l = newListener()
		s.addListener(l)
	}

	s.notify("Bidder1 100$!")
	s.notify("Bidder4 100$!")
}
