package main

import "fmt"

const greetingMsg string = "hello"

func HelloYou(name string) string {
	return fmt.Sprintf("%s, %s!", greetingMsg, name)
}
