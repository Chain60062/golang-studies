package main

import "fmt"

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}
