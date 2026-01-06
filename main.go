package main

import (
	"fmt"
	"reviewservice/config"
)

func main() {
	config.Load()
	fmt.Println("aksndlkansdl")
	config.SetupDb()
}