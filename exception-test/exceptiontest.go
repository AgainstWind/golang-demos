package main

import (
	"fmt"
	"sync"
)

func f()  {
	defer func() {
		fmt.Println("inner func defer...")
	}()

	defer func() {
		fmt.Println("ddd")
		if err:=recover();err != nil{
			fmt.Println(err)
		}
		fmt.Println("eeee")
	}()

	fmt.Println("A")

	panic(3)

	defer func() {
		fmt.Println("inner func defer..1")
	}()

	fmt.Println("b")
}


func main() {

	fmt.Println("Hello World!")

	defer func() {
		fmt.Println("Fuck-1")
	}()

	defer func() {
		fmt.Println("d")
		if err:=recover();err != nil{
			fmt.Println(err)
		}
		fmt.Println("e")
	}()

	defer func() {
		fmt.Println("Fuck-2")
	}()
	fmt.Println("Hello World! 2 ")

	f()
	fmt.Println("Hello World! 3")

	defer func() {
		fmt.Println("Fuck-3")
	}()

}
