package main

import (
    "fmt"
)

func main() {

    matched := false
    for n := 1; n < 101; n++ {
        if n%3 == 0 {
            matched = true
            fmt.Print("Fizz")
        }
        if n%5 == 0 {
            matched = true
            fmt.Print("Buzz")
        }
        if matched{
            fmt.Print("!")
        } else {
            fmt.Print(n)
        }
        fmt.Print("\n")
        matched = false
    }
}
