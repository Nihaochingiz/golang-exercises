package main

import "fmt"


func add(x int, y int) {
	total :=0
	total = x + y
	fmt.Println(total)
}



func main() {
	add(10,30)
}