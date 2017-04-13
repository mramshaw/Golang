// Some experiments & simple tests for 'panic' and 'defer'
//
// Go's panic/recover seems similiar to Java's try/catch/finally
//
//     should probably be reserved for the unexpected (exceptions)
//     as a better approach for so-called 'normal' exceptions (i.e.
//     expected errors as opposed to unexpected errors) is via the
//     'error' return value
//
// Go's error concept seems similiar to Java's exception concept
//
//     can define custom exceptions in Java, can define custom errors in Go
//
//         simply need to implement the Error() method on a custom struct
//
// 'defer' code gets pushed onto the return stack and will execute
// after the surrounding code block (i.e. curly braces) completes;
// but in LIFO order. Probably best to declare all 'defer' code at
// the start of the surrounding code block (unless it makes more
// sense programmatically or semantically to declare it later in
// the code block)
//
// Officially:
//
//     "A defer statement pushes a function call onto a list.
//     The list of saved calls is executed after the surrounding
//     function returns. Defer is commonly used to simplify functions
//     that perform various clean-up actions."
//
// 'recover' code needs to be 'defer' code as well and is unusual
// in that it gets executed even in a 'panic' whereas all other code
// gets bypassed in a panic
//
// Could be compiled but probably better to run with:
//
//     go run panic-defer.go
//
// Multiple 'return' statements within subroutines
//   should not have inconsistent behaviour.

package main

import "fmt"

func main() {
    fmt.Println("main() starting...")
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in main", r)
        }
    }()
    fmt.Println("Calling subRoutine()...")
    subRoutine()
    fmt.Println("Returned normally from subRoutine()")
    fmt.Println("main() completed")
}

func subRoutine() {
    fmt.Println("subRoutine() starting...")
//    defer func() {
//        if r := recover(); r != nil {
//            fmt.Println("Recovered in subRoutine", r)
//        }
//    }()
    fmt.Println("Calling recursiveSubRoutine...")
    recursiveSubRoutine(0)
    fmt.Println("Returned normally from recursiveSubRoutine")
}

func recursiveSubRoutine(i int) {
    defer func() {
        fmt.Println("Defer in recursiveSubRoutine", i)
    }()
    if i > 3 {
        fmt.Println("Panicking in recursiveSubRoutine!")
        panic(fmt.Sprintf("%v", i))
    }
    fmt.Println("Printing in recursiveSubRoutine", i)
    printIfNotEven(i)
    recursiveSubRoutine(i + 1)
}

func printIfNotEven(i int) {
    defer func() {
        fmt.Println("Defer in evenTest", i)
    }()
    if i % 2 == 0 {
        fmt.Println("Skipping even number!")
        return
    }
    fmt.Println("Printing in evenTest", i)
    return
}
