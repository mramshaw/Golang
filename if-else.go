package main

import "fmt"

func main() {

    // Here's a basic example.
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // You can have an `if` statement without an else.
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // A statement can precede conditionals; any variables
    // declared in this statement are available in all
    // branches.
    if num := 9; num%3 == 0 {
        fmt.Println(num, "is divisble by 3")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    // NOTA BENE: in the above 'if' the declared variables
    //   are ONLY available within the scope of code block.
    // Uncomment the following line to test:
//    fmt.Println("num = ", num)

    // If the above line is uncommented, the code will not
    // compile (as of GO 1.8):
    //
    //     ./if-else.go:33: undefined: num
}
