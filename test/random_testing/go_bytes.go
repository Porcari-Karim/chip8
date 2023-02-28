package main

import "fmt"

func main() {
	b := byte(0b00001111)
	b = b << 4
	fmt.Printf("%b \n", b)
	b = b + byte(0b00001111)
	fmt.Printf("%b \n", b)

}
