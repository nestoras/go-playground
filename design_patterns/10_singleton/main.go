package main

type config map[interface{}]interface{}

var value config

// Is it correct ?
// Are you familiar with this ?
func NewConfig() config {
	if value == nil {
		value = make(config)
	}
	return value
}

func main() {
	_ = NewConfig()
}
