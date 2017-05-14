# Golang
Experiments with Go

Go is a fun language but not without eccentricities.

It strikes a nice balance between the "raw metal" and having to know a zillion APIs in order to do anything.

One of the really nice features is that a core feature of the design is that Go 1.x releases should be "future-proof" (new point releases should not introduce incompatible APIs that require changing code - of course this does not apply to major releases such as Go 2.x or Go 3.x): https://golang.org/doc/go1compat

I'm enjoying it quite a lot and I am - on purpose - avoiding multi-threading.

[Multi-threading - and also recursion - can be quite tricky to debug and are definitely advanced topics.
 My personal opinion on recursion is that it should be avoided if at all possible - not always the case,
 I have certainly had to use recursion due to some unusual requirements; on the other hand multi-threading
 has many valid use-cases.]

As far as I can tell, Go was designed for Multi-threading (or at least Concurrency) - which is generally pretty tricky stuff. I'm looking forward to seeing how it handles multi-threading but first I want to get a good grasp
on the fundamentals.
