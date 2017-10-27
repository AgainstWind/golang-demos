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

	map1 := make(map[string]int) //key is string,value is int
	map1["key1"] = 2
	var map2 = map[string]int{
		"key1":1,
		"key2":32,
	}
	fmt.Println(map1,map2)
	valueTest,isExist := map2["key2"]
	fmt.Println(isExist,valueTest)

}

func f() *int {
	v := 1
	return &v
}
