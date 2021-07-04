package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type hindiBot struct{}

func main() {

	eb := englishBot{}

	hb := hindiBot{}

	printGreeting(eb)

	printGreeting(hb)

}

func (englishBot) getGreeting() string {
	return "Hi there !!"
}

func (hindiBot) getGreeting() string {
	return "Namaste !!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
