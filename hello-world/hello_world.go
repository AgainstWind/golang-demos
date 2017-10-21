package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello Golang!\n")
	x := 1
	p := &x
	fmt.Printf(":%d\n", *p)
	fmt.Printf(":%d\n", p)
	fmt.Printf(":%d\n", x)
	var d *int
	d=f()
	fmt.Printf(":%d\n", d)
	fmt.Printf(":%d\n", f())
	fmt.Printf(":%d\n", *d)
}



func f() *int{
	v :=1
	return &v
}
