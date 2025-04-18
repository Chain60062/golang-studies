package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Loop", i)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }
	// 	samples := []string{"hello", "apple_π!"}

	// outer:
	// 	for _, sample := range samples {
	// 		for i, r := range sample {
	// 			fmt.Println(i, r, string(r))

	// 			if r == 'l' {
	// 				continue outer
	// 			}
	// 		}
	// 	}

	// evenVals := []int{2, 4, 6, 8, 10}

	// for i, v := range evenVals {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	if i == len(evenVals)-1 {
	// 		break
	// 	}
	// 	fmt.Println(i, v)
	// }

	// for i := 1; i < len(evenVals)-1; i++ {
	// 	fmt.Println(i, evenVals[i])
	// }

	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

}
