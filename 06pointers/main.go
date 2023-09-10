package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers !")
	var ptr *int
	fmt.Printf("Value of pointer is %v", ptr) // When not initialized, the value of pointer is always nil
	myNumber := 20

	var numPtr = &myNumber
	fmt.Println("Value of the pointer is ", numPtr)
	fmt.Println("Actual Value of the pointer is ", *numPtr)

	*numPtr = *numPtr * 2
	fmt.Println("Actual Value of the pointer is ", *numPtr)
	fmt.Println("Actual Value of the pointer is ", myNumber)

}
