package main

import (
	"fmt"
)

const defaultGreeting string = "Hello, World!"
const greetingMsg string = "Hello"

func main() {
	fmt.Println(helloWorldAgain("mateen"))
}

func helloWorldAgain(name string) string {
	if name != "" {
		return fmt.Sprintf("%s, %s", greetingMsg, name)
	}
	return defaultGreeting
}
