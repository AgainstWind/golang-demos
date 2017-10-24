package main

import (
	"fmt"
	"math/cmplx"
	"strings"
	"math"
)

func main() {
	fmt.Printf("Hello Golang!\n")
	x := 1
	p := &x
	fmt.Printf(":%d\n", *p)
	fmt.Printf(":%d\n", p)
	fmt.Printf(":%d\n", x)
	var d *int
	d = f()
	fmt.Printf(":%d\n", d)
	fmt.Printf(":%d\n", f())
	fmt.Printf(":%d\n", *d)

	fmt.Println(cmplx.Sqrt(-1))

	var testString string = `hsfgys\nfsfs""`
	var runeTest rune = 'd'

	fmt.Println(testString)
	fmt.Println(runeTest)
	fmt.Println(strings.Contains(testString,"fg"))
	var a = math.Pi
	fmt.Println(a)

}

func f() *int {
	v := 1
	return &v
}
