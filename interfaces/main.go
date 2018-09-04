package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreetingSB(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSB(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
