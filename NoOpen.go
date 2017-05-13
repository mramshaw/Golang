package main

import (
    "fmt"
    "os"
)

func main() {

    _, err := os.Open("NoSuchFile")                     // File does not exist

    if err != nil {

        fmt.Printf("%s\n", err)                         // generic error string

        fmt.Printf("%#v\n", err)                        // specific error details

        if pe, ok := err.(*os.PathError); ok {          // specific type assertion
            fmt.Printf("PathError asserted: %#v\n", pe)
        }
    }
}
