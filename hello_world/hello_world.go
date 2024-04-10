package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
    fmt.Println(Hello("Thomas"))
    }

func Hello(s string) string {
    if s == "" {
      return "Hello, World"
    }
    return englishHelloPrefix + s
}