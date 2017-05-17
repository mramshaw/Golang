// Some gotchas to be aware of when creating Go routines.
//
// 1) Pass-by-reference versus Pass-by-copy can be an issue.
//
// 2) The order in which Go routines actually execute can be
//    random (non-deterministic).
//
// 3) Go routines can be stranded (leak) which is considered
//    bad practice. Adding a panic before the end of the 'main'
//    Go routine (recommended while testing) does not seem to
//    provide a list of all stranded Go routines as expected.
//
// @ Martin Ramshaw, May 2017 (mramshaw@alumni.concordia.ca)

package main

import (
	"fmt"
	"runtime"
)

const (
	version         = "1.0"
	defaultNumber   = 5
)

func main() {

	fmt.Printf("\n== Gotchas %s (runtime: %s - CPUs: %d) == Ctrl-C to quit!\n", version, runtime.Version(), runtime.NumCPU())
	fmt.Printf("\n")
	fmt.Printf("Number of Go routines (pre-loop) : = %d\n", runtime.NumGoroutine())

	// Order in which scheduler runs goroutines is non-deterministic
	for i := 0; i < defaultNumber; i++ {
		go printNumbers(i, &i)		// value of &i will probably be defaultNumber (5)
	}
	fmt.Printf("Number of Go routines (post-loop): = %d\n", runtime.NumGoroutine())

	// uncomment the next line for leaking Go routines
//	panic(fmt.Sprintf("Probably leaking some Go routines"))

	for runtime.NumGoroutine() > 1 {}	// comment this line to leak Go routines
}

func printNumbers(a int, b *int) {

	fmt.Printf("a = %d, b = %d\n", a, *b)
}
