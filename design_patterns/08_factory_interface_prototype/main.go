package main

import "fmt"

type Services struct {
	serviceA ServiceA
}

type ServicesConfig func(*Services) error

func NewServices(cfgs ...ServicesConfig) (*Services, error) {
	var s Services
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

type ServiceA interface {
	GetTitle() string
	GetCategory() string
	GetTitleAndCategory() string
}

func (s *serviceA) GetTitle() string {
	return s.title
}

func (s *serviceA) GetCategory() string {
	return s.title
}

func (s *serviceA) GetTitleAndCategory() string {
	return "title: " + s.title + " category: " + s.category
}

func NewServiceA(title, category string) ServiceA {
	return &serviceA{
		title:    title,
		category: category,
	}
}

// It creates an unused instance of UserService, that will signal us if the UserService interface is not implemented
// correctly. The line isn't 100% necessary and is more like a test that happens to lead to a compiler errors when it
// fails
var _ ServiceA = &serviceA{}

type serviceA struct {
	title    string
	category string
}

func WithServiceA(title, category string) ServicesConfig {
	return func(s *Services) error {
		s.serviceA = NewServiceA(title, category)
		return nil
	}
}

func main() {
	services, err := NewServices(WithServiceA("golang", "course"))
	if err != nil {
		panic(err)
	}
	title := services.serviceA.GetTitle()
	fmt.Println(title)

	titleCategory := services.serviceA.GetTitleAndCategory()
	fmt.Println(titleCategory)
}
