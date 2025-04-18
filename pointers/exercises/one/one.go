package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstname, lastname string, age int) Person {
	return Person{firstname, lastname, age}
}

func MakePersonPointer(firstname, lastname string, age int) *Person {
	return &Person{firstname, lastname, age}
}
func main() {
	person1 := MakePerson("Vinicius", "Miranda", 22) //person1 escapes to heap because of the call to fmt.Println
	fmt.Println(person1)
	person2 := MakePersonPointer("Vinicius", "Miranda", 22)
	fmt.Println(person2)
}

/*
 The &Person{} returned from MakePersonPointer escapes to the heap. Any time a pointer is returned from a function, the pointer is returned on the stack, but the value it points to must be stored on the heap.

 Surprisingly, it also says that p escapes to the heap. The reason for this is that it is passed into fmt.Println. This is because the parameter to fmt.Println are ...any. The current Go compiler moves to the heap any value that is passed in to a function via a parameter that is of an interface type.
*/
