package main

import "fmt"

var helloWorldMessage string = "Mateen has started learning more about TDD!"

func main() {
	fmt.Println(hello())
}

func hello() string {
	return helloWorldMessage
}
