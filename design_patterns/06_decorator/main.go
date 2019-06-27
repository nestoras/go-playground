package main

import "fmt"

type Window interface {
	Draw()
	GetDescription() string
}

type SimpleWindow struct{}

func (sw *SimpleWindow) Draw() {
	//draw a simple window
}
func (sw *SimpleWindow) GetDescription() string {
	return "a simple window"
}

type WindowDecorator struct {
	window Window
}

func NewWindowDecorator(window Window) *WindowDecorator {
	return &WindowDecorator{window}
}

func (wd *WindowDecorator) Draw() {
	wd.window.Draw()
}
func (wd *WindowDecorator) GetDescription() string {
	return wd.window.GetDescription()
}

type VerticalScrollBars struct {
	*WindowDecorator
}

func NewVerticalScrollBars(window Window) *VerticalScrollBars {
	vsb := VerticalScrollBars{}
	vsb.WindowDecorator = NewWindowDecorator(window)
	return &vsb
}

func (vsb *VerticalScrollBars) Draw() {
	vsb.drawVerticalScrollBar()
	vsb.window.Draw()
}

func (vsb *VerticalScrollBars) drawVerticalScrollBar() {
	fmt.Println("vertical scrollbar")
}

func (vsb *VerticalScrollBars) GetDescription() string {
	return vsb.window.GetDescription() + ", including vertical scrollbars"
}

type HorizontalScrollBars struct {
	*WindowDecorator
}

func NewHorizontalScrollBars(window Window) *HorizontalScrollBars {
	hsb := HorizontalScrollBars{}
	hsb.WindowDecorator = NewWindowDecorator(window)
	return &hsb
}

func (hsb *HorizontalScrollBars) Draw() {
	hsb.drawVerticalScrollBar()
	hsb.window.Draw()
}

func (hsb *HorizontalScrollBars) drawVerticalScrollBar() {
	fmt.Println("horizontal scrollbar")
}

func (hsb *HorizontalScrollBars) GetDescription() string {
	return hsb.window.GetDescription() + ", including horizontal scrollbars"
}

func main() {
	var window = NewHorizontalScrollBars(
		NewVerticalScrollBars(
			new(SimpleWindow)))
	fmt.Println(window.GetDescription())
}
