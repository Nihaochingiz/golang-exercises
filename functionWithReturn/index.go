package main

import "fmt"


func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)

	area = l*b
	return
}


func main() {
	fmt.Println("Area: ", rectangle(20, 30))
}