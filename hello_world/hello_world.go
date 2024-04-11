package main

import "fmt"

const (
    french = "French"
    spanish = "Spanish"

    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix = "Bonjour, "
)

func main() {
    fmt.Println(Hello("Thomas", "Spanish"))
    }

func Hello(s string, l string) string {
    if s == "" {
      return "Hello, World"
    }

    return greeting(l) + s
}

func greeting(l string) (prefix string) {
    switch l {
    case spanish:
        prefix = spanishHelloPrefix
    case french:
        prefix =  frenchHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
   return
}