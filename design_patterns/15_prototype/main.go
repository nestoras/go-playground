package main

import "fmt"
import "errors"

const (
	YELLOW = 1
	BLACK  = 2
	BLUE   = 3
)

type Getter interface {
	GetShirt() string
}

type Shirt struct {
	Size  string
	Color string
}

func (s *Shirt) GetShirt() string {
	return fmt.Sprintf("Shirt size: '%s', Color: '%s'.", s.Size, s.Color)
}

var yellowShirt *Shirt = &Shirt{"S", "Yellow"}
var blackShirt *Shirt = &Shirt{"M", "Black"}
var blueShirt *Shirt = &Shirt{"L", "Blue"}

func NewShirt(t int) (Getter, error) {
	switch t {
	case YELLOW:
		return yellowShirt, nil
	case BLACK:
		return blackShirt, nil
	case BLUE:
		return blueShirt, nil
	default:
		return nil, errors.New(fmt.Sprintf("Wrong id %d.", t))
	}
}

func main() {
	yellow, err := NewShirt(YELLOW)
	if err != nil {
		panic(err)
	}

	details := yellow.GetShirt()
	fmt.Println(details)
}
