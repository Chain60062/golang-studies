package main

import "fmt"

func main() {
	fmt.Println(add(2, 3)) // Error: undefined add
}

func add(a int, b int) int {
	return a + b
}
