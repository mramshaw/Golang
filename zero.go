package main

import (
    "fmt"
    "reflect"
)

type Vertex struct {
    X, Y int
}

// ======================

type customError struct {
    arg  int
    desc string
}

func (ce *customError) Error() string {
    return fmt.Sprintf("%d - %s", ce.arg, ce.desc)
}

// ======================

func main() {

    var b    bool

    var i    int
    var u    uint
    var i8   int
    var u8   uint
    var i16  int
    var u16  uint
    var i32  int32
    var u32  uint32
    var i64  int64
    var u64  uint64

    var f32  float32
    var f64  float64

    var c64  complex64
    var c128 complex128

    var s string

    var v Vertex

    var ce customError

    fmt.Printf("\n")
    fmt.Printf("Booleans\n")
    fmt.Printf("========\n")
    fmt.Printf("\n")
    fmt.Printf("bool is initialized to: %v, type is %T\n", b, b)

    fmt.Printf("\n")
    fmt.Printf("Integers\n")
    fmt.Printf("========\n")
    fmt.Printf("\n")
    fmt.Printf("int   is initialized to: %v, type is %T\n", i,   i)
    fmt.Printf("int8  is initialized to: %v, type is %T\n", i8,  i8)
    fmt.Printf("int16 is initialized to: %v, type is %T\n", i16, i16)
    fmt.Printf("\n")
    fmt.Printf("uint   is initialized to: %v, type is %T\n", u,   u)
    fmt.Printf("uint8  is initialized to: %v, type is %T\n", u8,  u8)
    fmt.Printf("uint16 is initialized to: %v, type is %T\n", u16, u16)
    fmt.Printf("\n")
    fmt.Printf("int32  is initialized to: %v, type is %T\n", i32, i32)
    fmt.Printf("uint32 is initialized to: %v, type is %T\n", u32, u32)
    fmt.Printf("\n")
    fmt.Printf("int64  is initialized to: %v, type is %T\n", i64, i64)
    fmt.Printf("uint64 is initialized to: %v, type is %T\n", u64, u64)
    fmt.Printf("\n")
    fmt.Printf(" ** int should be 32 bits on 32-bit systems and 64 bits on 64-bit systems\n")
    fmt.Printf("    if unsure, probably best to specify rather than leave it to the compiler\n")

    fmt.Printf("\n")
    fmt.Printf("Floats\n")
    fmt.Printf("======\n")
    fmt.Printf("\n")
    fmt.Printf("float32 is initialized to: %v, type is %T\n", f32, f32)
    fmt.Printf("float64 is initialized to: %v, type is %T\n", f64, f64)

    fmt.Printf("\n")
    fmt.Printf("Complex\n")
    fmt.Printf("=======\n")
    fmt.Printf("\n")
    fmt.Printf("complex64  is initialized to: %v, type is %T\n", c64,  c64)
    fmt.Printf("complex128 is initialized to: %v, type is %T\n", c128, c128)

    fmt.Printf("\n")
    fmt.Printf("Strings\n")
    fmt.Printf("=======\n")
    fmt.Printf("\n")
    fmt.Printf("string is initialized to: %q, type is %T\n", s, s)

    fmt.Printf("\n")
    fmt.Printf("Structs (using Reflection)\n")
    fmt.Printf("==========================\n")
    fmt.Printf("\n")
    fmt.Printf("%#v\n", v)
    fmt.Printf("\n")
    fmt.Printf("struct is initialized to: %v, type is %T\n", v, v)
    fmt.Println("    reflect.TypeOf(v) = ", reflect.TypeOf(v))
    fmt.Println("    reflect.ValueOf(v) = ", reflect.ValueOf(v))
    fmt.Println("    reflect.ValueOf(v).Kind() = ", reflect.ValueOf(v).Kind())

    fmt.Printf("\n")
    fmt.Printf("Errors (using Reflection)\n")
    fmt.Printf("=========================\n")
    fmt.Printf("\n")
    fmt.Printf("%#v\n", ce)
    fmt.Printf("\n")
    fmt.Printf("error is initialized to: %v, type is %T\n", ce, ce)
    fmt.Println("    reflect.TypeOf(ce) = ", reflect.TypeOf(ce))
    fmt.Println("    reflect.ValueOf(ce) = ", reflect.ValueOf(ce))
    fmt.Println("    reflect.ValueOf(ce).Kind() = ", reflect.ValueOf(ce).Kind())

    fmt.Printf("\n")
    fmt.Printf("Use of Custom Error (through casting)\n")
    fmt.Printf("=====================================\n")
    fmt.Printf("\n")
    rc, err := evenOdd(4)
    if ee, ok := err.(*customError); ok {
        fmt.Println("Custom Error caught, ", err)
        fmt.Println("    Custom Error: ", ee.arg)
        fmt.Println("    Custom Error: ", ee.desc)
    }
    fmt.Println("Return code = ", rc)

    fmt.Printf("\n")
    fmt.Printf("Non-Use of Custom Error (through casting)\n")
    fmt.Printf("=========================================\n")
    fmt.Printf("\n")
    rc, err = evenOdd(5)
    if ee, ok := err.(*customError); ok {
        fmt.Println("Custom Error caught, ", err)
        fmt.Println("    Custom Error: ", ee.arg)
        fmt.Println("    Custom Error: ", ee.desc)
    } else {
        fmt.Println("No Error!")
    }
    fmt.Println("Return code = ", rc)

    var chBool   chan bool
    var chInt    chan int
    var chByte   chan byte
    var chString chan string
    var chIface  chan interface{}

    fmt.Printf("\n")
    fmt.Printf("Channels\n")
    fmt.Printf("========\n")
    fmt.Printf("\n")
    fmt.Printf("%#v\n", chBool)
    fmt.Printf("%#v\n", chInt)
    fmt.Printf("%#v\n", chByte)
    fmt.Printf("%#v\n", chString)
    fmt.Printf("%#v\n", chIface)
    fmt.Printf("\n")
    fmt.Printf("bool      channel is initialized to: %[01]v, type is %[01]T\n", chBool)
    fmt.Printf("int       channel is initialized to: %[01]v, type is %[01]T\n", chInt)
    fmt.Printf("byte      channel is initialized to: %[01]v, type is %[01]T\n", chByte)
    fmt.Printf("string    channel is initialized to: %[01]v, type is %[01]T\n", chString)
    fmt.Printf("interface channel is initialized to: %[01]v, type is %[01]T\n", chIface)

    chBool   = make(chan bool)
    chInt    = make(chan int)
    chByte   = make(chan byte)
    chString = make(chan string)
    chIface  = make(chan interface{})

    fmt.Printf("\n")
    fmt.Printf("bool      channel (after make) is: %[01]v, type is %[01]T\n", chBool)
    fmt.Printf("int       channel (after make) is: %[01]v, type is %[01]T\n", chInt)
    fmt.Printf("byte      channel (after make) is: %[01]v, type is %[01]T\n", chByte)
    fmt.Printf("string    channel (after make) is: %[01]v, type is %[01]T\n", chString)
    fmt.Printf("interface channel (after make) is: %[01]v, type is %[01]T\n", chIface)

    close(chBool)
    close(chInt)
    close(chByte)
    close(chString)
    close(chIface)

    fmt.Printf("\n")
    fmt.Printf("bool      channel (after close): %[01]v, type is %[01]T\n", chBool)
    fmt.Printf("int       channel (after close): %[01]v, type is %[01]T\n", chInt)
    fmt.Printf("byte      channel (after close): %[01]v, type is %[01]T\n", chByte)
    fmt.Printf("string    channel (after close): %[01]v, type is %[01]T\n", chString)
    fmt.Printf("interface channel (after close): %[01]v, type is %[01]T\n", chIface)

    fmt.Printf("\n")
    fmt.Printf("bool      channel (read after close) gives: %v\n", <-chBool)
    fmt.Printf("int       channel (read after close) gives: %v\n", <-chInt)
    fmt.Printf("byte      channel (read after close) gives: %v\n", <-chByte)
    fmt.Printf("string    channel (read after close) gives: %q\n", <-chString)
    fmt.Printf("interface channel (read after close) gives: %v\n", <-chIface)
}

// Returns -1 and customError if Even, normal return if Odd
func evenOdd(arg int) (int, error) {

    if arg % 2 == 0 {
        return -1, &customError{arg, "Even number"}
    }

    return arg, nil
}

