package main

import (
	"fmt"
	"go-playground/general/01_config/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)
}
