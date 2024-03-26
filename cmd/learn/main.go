package main

import "fmt"

const (
	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return prefix(language) + name
}

func prefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Ethan", "english"))
}
