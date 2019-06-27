package main

import "fmt"

type Product struct {
	Type         string
	Manufacturer string
	Model        string
}

type Builder interface {
	SetType() Builder
	SetManufacturer() Builder
	SetModel() Builder
	GetInfo() Product
}

type BuilderDirector struct {
	builder Builder
}

func (bd *BuilderDirector) SetBuilder(b Builder) {
	bd.builder = b
}

func (bd *BuilderDirector) Construct() Product {
	bd.builder.SetType().SetManufacturer().SetModel()
	return bd.builder.GetInfo()
}

func (bd *BuilderDirector) GetProduct() {
	gadget := bd.builder.GetInfo()
	fmt.Printf("Type of product: %s \n", gadget.Type)
	fmt.Printf("Manufacturer: %s \n", gadget.Manufacturer)
	fmt.Printf("Model: %s \n", gadget.Model)

}

type MobilePhone struct {
	product Product
}

func (m *MobilePhone) SetType() Builder {
	m.product.Type = "mobile"
	return m
}

func (m *MobilePhone) SetManufacturer() Builder {
	m.product.Manufacturer = "Apple"
	return m
}

func (m *MobilePhone) SetModel() Builder {
	m.product.Model = "iPhone X"
	return m
}

func (m *MobilePhone) GetInfo() Product {
	return m.product
}

type Television struct {
	product Product
}

func (t *Television) SetType() Builder {
	t.product.Type = "television"
	return t
}

func (t *Television) SetManufacturer() Builder {
	t.product.Manufacturer = "LG"
	return t
}

func (t *Television) SetModel() Builder {
	t.product.Model = "X100"
	return t
}

func (t *Television) GetInfo() Product {
	return t.product
}

func main() {
	builderDirector := BuilderDirector{}

	mobile := &MobilePhone{}
	builderDirector.SetBuilder(mobile)
	builderDirector.Construct()
	builderDirector.GetProduct()

	television := &Television{}
	builderDirector.SetBuilder(television)
	builderDirector.Construct()
	builderDirector.GetProduct()
}
