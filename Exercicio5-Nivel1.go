package main

import "fmt"

type inteiro int

var y int
var x inteiro

func main() {
  
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)

	y = int(x)

	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
  
}
