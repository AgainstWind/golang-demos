package main

import "net/http/httputil"
import "fmt"

func log(msg string){
	fmt.Printf(msg)
}

func main()  {
	go log("test")
}