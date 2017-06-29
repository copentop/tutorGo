package main

import "fmt"

type Dog struct {
    Name string
    Type string
}
func addressTest(d *Dog) {
    a := &Dog{"another cute dog", "another type"}
    d = a
}

func main() {
	var a, b Dog;

	a.Name = "5"
	a.Type = "6"

	addressTest(&a)
	addressTest(&b)

	fmt.Println("Dog ", a)
	fmt.Println("Another Dog ", b)

}
