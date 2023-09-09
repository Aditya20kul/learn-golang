package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Welcome to User Input"
	fmt.Print(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a rating for the food out of 5: ")

	// comma ok syntax || err ok
	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for rating us %s stars !", input)
}
