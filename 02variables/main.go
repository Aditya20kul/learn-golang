package main

import "fmt"

var JwtToken string = "efdxvyw562374d2w623des26" // When the First letter of a variable is capital, It is considered to be a public variable 

func main() {
	fmt.Println("================ Variables ===================")
	var username string = "Aditya"
	var isLoggedIn bool = false
	fmt.Printf("Type of username = %T \n Type of isLoggedIn = %T\n", username, isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type of smallVal - %T\n", smallVal)

	var smallFloat float64 = 255.222222222232321232143
	fmt.Println(smallFloat)
	fmt.Printf("Type of smallVal - %T\n", smallFloat)

	// Another way of declaring variables
	var website = "Aditya K"
	println(website)

	// no var style - This type of declaration is only allowed inside of a method (Meaning, It cannot be used when doing global declarations)
	numberOfUsers := 300000
	println(numberOfUsers)
}
