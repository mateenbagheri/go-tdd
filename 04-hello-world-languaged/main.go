package main

import (
	"fmt"
)

const defaultGreeting string = "Hello, World!"
const defaultGreetingSpanish string = "Hola, World!"
const greetingMsg string = "Hello"
const greetingMsgSpanish string= "Hola"

func main() {
	fmt.Println(helloWordLanguaged("mateen", "Spanish"))
}

func helloWordLanguaged(name, language string) string {
	if language == "Spanish" {
		if name == "" {
			return defaultGreetingSpanish
		}
		return fmt.Sprintf("%s, %s", greetingMsgSpanish, name)
	}
	if name != "" {
		return fmt.Sprintf("%s, %s", greetingMsg, name)
	}
	return defaultGreeting
}
