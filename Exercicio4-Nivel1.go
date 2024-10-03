package main

import "fmt"

type inteiro int
var x inteiro

func main() {

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
  
}
