package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	firstTwo := greetings[:2]
	three := greetings[1:4]
	lastTwo := greetings[3:5]

	fmt.Println(firstTwo)
	fmt.Println(three)
	fmt.Println(lastTwo)
}
