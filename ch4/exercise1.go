package main

import "fmt"

func main() {
	//Exercise 1
	// var numbers []int
	// for i := 0; i < 100; i++ {
	// 	numbers = append(numbers, rand.Intn(100))
	// }

	// fmt.Println(numbers, len(numbers))

	//Exercise 2 (Depends on 1)
	// for _, num := range numbers {
	// 	switch {
	// 	case num%2 == 0 && num%3 == 0:
	// 		fmt.Println(num, "Six")
	// 	case num%2 == 0:
	// 		fmt.Println(num, "Two")
	// 	case num%3 == 0:
	// 		fmt.Println(num, "Three")
	// 	default:
	// 		fmt.Println("Nevermind")
	// 	}
	// }

	//Exercise 3
	var total int

	for i := 0; i < 10; i++ {
		total = total + 1
		fmt.Println(total)
	}

}
