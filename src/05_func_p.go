package main

import (
    "fmt"
)

func main() {
    defer func() {
        defer func() {
            fmt.Println("6: def in rec:", recover())
        }()
    }()
    defer func() {
        func() {
            fmt.Println("5: def in rec:", recover())
        }()
    }()
	
    // -------
    func() {
        defer func() {
            fmt.Println("1: def in rec:", recover())
        }()

	fmt.Println("f 1")
    }()
    func() {
        defer fmt.Println("2: def rec:", recover())

	fmt.Println("f 2")
    }()

    func() {
        fmt.Println("3: rec:", recover())
	
	fmt.Println("f 3")
    }()

    fmt.Println("4: rec rec:", recover())

    panic("err")


    defer func() {
        fmt.Println("0: rec:", recover()) // never go here
    }()
}
