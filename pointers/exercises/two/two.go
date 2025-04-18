package main

import (
	"errors"
	"fmt"
	"log"
)

func UpdateSlice(slice []string, text string) error {
	if len(slice) == 0 {
		return errors.New("slice vazio")
	}
	slice[len(slice)-1] = text
	fmt.Println(slice)
	return nil
}
func GrowSlice(slice []string, text string) error {
	if len(slice) == 0 {
		return errors.New("slice vazio")
	}
	slice = append(slice, text) //does nothing
	return nil
}

func main() {
	mySlice := make([]string, 100)
	mySlice[0] = "one"
	mySlice[1] = "two"

	err := UpdateSlice(mySlice, "one hundred")

	if err != nil {
		log.Fatal(err)
	}

	mySlice2 := []string{"one", "two"}
	fmt.Println(mySlice2)
	err = GrowSlice(mySlice2, "three")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mySlice2)
}
