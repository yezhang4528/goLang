package main

import "fmt"

type bot interface {
    getGreeting() string
}

type EnglishGreeting struct {}
type SpanishGreeting struct {}

func main() {
    eb := EnglishGreeting{}
    sb := SpanishGreeting{}

    printGreeting(eb)
    printGreeting(sb)
}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

func (EnglishGreeting) getGreeting() string {
    return "Hi there!"
}

func (SpanishGreeting) getGreeting() string {
    return "Hola!"
}
