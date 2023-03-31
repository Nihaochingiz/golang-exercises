package main
import "fmt"

func main() {
	variadicExample("red", " blue",
"green", "yellow","orange", "white")
}


func variadicExample( s ...string) {
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[3])
	fmt.Println(s[5])
}