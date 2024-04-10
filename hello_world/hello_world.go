package main

import "fmt"

func main() {
    fmt.Println(Hello("Thomas"))
    }

func Hello(s string) string {
    return fmt.Sprintf("Hello, %s", s)
}