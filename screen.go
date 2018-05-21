package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	err := errors.New("new error")

	fmt.Println("err = ", err)

    // This is a nice example of scoping in Go.
    // The 'err' field created in the next line
    // will override (but not replace) the previous
    // value of 'err'. The scope of the created 'err'
    // field is the lifetime of the 'if' statement.
	if _, err := os.Stat("does_not_exist"); err != nil {
		fmt.Println("err = ", err)
	}

	fmt.Println("err = ", err)
}
