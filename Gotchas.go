// Some gotchas to be aware of when creating Go routines.
//
// 1) Pass-by-reference versus Pass-by-copy can be an issue.
//
// 2) The order in which Go routines actually execute can be
//    random (non-deterministic).
//
// 3) Go routines can be stranded (leak) which is considered
//    bad practice. Adding a panic before the end of the 'main'
//    Go routine (recommended while testing) will provide a
//    list of all stranded Go routines if the runtime/debug
//    package's SetTraceback function is called as follows:
//
//        debug.SetTraceback("all")
//
//    A better way to do this is one of the following:
//
//        $ GOTRACEBACK="all" go run Gotchas.go
//
//        $ GOTRACEBACK=all go run Gotchas.go
//
//        $ GOTRACEBACK=1 go run Gotchas.go
//
//    [The results are identical and avoid having to
//     import the runtime/debug package. These can also
//     be specified at run time as needed - but since the
//     panic statement needs to be either commented or
//     uncommented as necessary the actual utility
//     of this when testing is still problematic.] 
//
//    The default value of Traceback is single (only show
//    the current Go routine).
//
// @ Martin Ramshaw, May 2017 (mramshaw@alumni.concordia.ca)

package main

import (
	"fmt"
	"runtime"
//	"runtime/debug"
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

	// uncomment the next 2 lines for leaking Go routines
//	debug.SetTraceback("all")
//	panic(fmt.Sprintf("Probably leaking some Go routines"))

	for runtime.NumGoroutine() > 1 {}	// comment this line to leak Go routines
}

func printNumbers(a int, b *int) {

	fmt.Printf("a = %d, b = %d\n", a, *b)
}
