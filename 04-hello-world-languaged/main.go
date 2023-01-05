package main

import (
	"fmt"
)

const engPrefix string = "Hello"
const frPrefix string = "Bonjour"
const spPrefix string = "Hola"

const defaultSpWorld string = "Mundo"
const defaultFrWorld string = "Monde"
const defaultEngWorld string = "World"

func main() {
	fmt.Println(helloWordLanguaged("mateen", "Spanish"))
}

func helloWordLanguaged(name, language string) string {
	prefix := prefixGenerator(language)
	if name == "" {
		world := worldWordGenerator(language)
		return fmt.Sprintf("%s, %s", prefix, world)
	}
	return fmt.Sprintf("%s, %s", prefix, name)
}

func worldWordGenerator(language string) string {
	var world string
	switch language {
	case "Spanish":
		world = defaultSpWorld
	case "French":
		world = defaultFrWorld
	case "English":
		world = defaultEngWorld
	default:
		world = defaultEngWorld
	}
	return world
}

func prefixGenerator(language string) string {
	var prefix string
	switch language {
	case "Spanish":
		prefix = spPrefix
	case "French":
		prefix = frPrefix
	case "English":
		prefix = engPrefix
	default:
		prefix = engPrefix
	}
	return prefix
}
