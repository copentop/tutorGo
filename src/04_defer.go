package main

import "fmt"


func main() {
	var a = 0;

	a = f()

	fmt.Println(a)
	a = f2()

	fmt.Println(a)
}

func f() (r int) {
     t := 5
     defer func() {
       t = t + 5
     }()
     return t
}

func f2() (r int) {
    defer func(r int) {
          r = r + 5
    }(r)
    return 1
}
