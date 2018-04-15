# Golang
Experiments with Go

[![GoDoc](https://godoc.org/github.com/mramshaw/Golang?status.svg)](https://godoc.org/github.com/mramshaw/Golang)

Go is a fun language but not without eccentricities.

It strikes a nice balance between the "raw metal" (like 'C') and having to know a zillion APIs in order to do anything.

One really nice thing is that a core feature of the design is that Go 1.x releases should be "future-proof" (new point
releases should not introduce incompatible APIs that require changing code - of course this does not apply to major
releases such as Go 2.x or Go 3.x): https://golang.org/doc/go1compat

Go has the concept of Concurrency (lightweight processes similiar to 'green' threads) which is not quite the same thing
as Multi-threading (threads are an OS concept and generally more limited in number). Go will use threads (say when
calling out to 'C' functions) but a blocking thread is a bit of an issue in Go whereas a blocking Go routine is not as
much of an issue (the Go scheduler will follow the Apache 2.4 model and conceptually shuffles the blocking Go routine
off onto a background process to be reactivated when the blocking event occurs).

My personal opinion on recursion is that it should be avoided if possible - however it is less of an issue in Go compared
to other languages. Stack Overflows are rare (if not impossible) in Go due to the design of the Go stack. While an OS
thread stack may be as much as 2 MB a typical Go routine stack starts at around 2 KB and can grow up to 1 GB (The Go
Programming Language, Donovan & Kernighan, page 280).

As far as I can tell, Go was designed for Concurrency.
