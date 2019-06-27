package main

import (
	"errors"
	"fmt"
)

type Drawer interface {
	DrawRectangle(w, h int)
}

type GreenRectangleDrawer struct{}

func (d GreenRectangleDrawer) DrawRectangle(w, h int) {
	fmt.Println("Green rectangle Width: %d, Height: %d", w, h)
}

type RedRectangleDrawer struct{}

func (d RedRectangleDrawer) DrawRectangle(w, h int) {
	fmt.Println("Red rectangle Width: %d, Height: %d", w, h)
}

type Rectangle struct {
	Width, Height int
	drawer        Drawer
}

type Car struct {
	Model        string
	Manufacturer string
	drawer       Drawer
}

func (d Rectangle) Draw() error {
	if d.drawer == nil {
		return errors.New("Drawer not initialized.")
	}

	d.drawer.DrawRectangle(d.Width, d.Height)
	return nil
}

func main() {
	green := GreenRectangleDrawer{}
	rectangle := Rectangle{
		Width:  200,
		Height: 200,
		drawer: green,
	}

	err := rectangle.Draw()
	must(err)

	red := RedRectangleDrawer{}
	rectangle2 := Rectangle{
		Width:  200,
		Height: 200,
		drawer: red,
	}
	err = rectangle2.Draw()
	must(err)

	rectangle3 := Rectangle{
		Width:  200,
		Height: 200,
		//drawer: red,
	}
	err = rectangle3.Draw()
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
