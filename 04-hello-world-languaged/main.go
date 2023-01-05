package main

import (
	"fmt"
)

const defaultGreeting string = "Hello, World!"
const defaultGreetingSpanish string = "Hola, World!"
const greetingMsg string = "Hello"
const greetingMsgSpanish string= "Hola"
const defaultGreetingFrench string = "Bonjour, World!"
const greetingMsgFrench string = "Bonjour"

func main() {
	fmt.Println(helloWordLanguaged("mateen", "Spanish"))
}

func helloWordLanguaged(name, language string) string {
	if language == "Spanish" {
		if name == "" {
			return defaultGreetingSpanish
		}
		return fmt.Sprintf("%s, %s", greetingMsgSpanish, name)
	} else if language == "French" {
		if name == "" {
			return defaultGreetingFrench
		}
		return fmt.Sprintf("%s, %s", greetingMsgFrench, name)
	}
	if name != "" {
		return fmt.Sprintf("%s, %s", greetingMsg, name)
	}
	return defaultGreeting
}
