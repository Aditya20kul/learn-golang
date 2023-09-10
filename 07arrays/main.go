package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays!")
	var fruitList [3]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Orange"

	fmt.Println(fruitList)
	fmt.Printf("Length of array - %v\n", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "mushroom"}
	fmt.Printf("Vegie List is - %v", vegList)
}
