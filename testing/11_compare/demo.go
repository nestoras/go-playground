package compare

func Square(i int) int {
	return i * i
}

type Dog struct {
	Name string
	Age  int
}

type DogWithFn struct {
	Name string
	Age  int
	Fn   func()
}