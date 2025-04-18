package main

import "fmt"

func main() {
	message := "Hi 👩🏻 and 👨🏻"

	var messageRunes []rune = []rune(message)

	fmt.Println(string(messageRunes[3]))
}
