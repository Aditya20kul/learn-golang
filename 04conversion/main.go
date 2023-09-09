package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizzaa App !")
	fmt.Printf("Please Rate our Pizza between 1 and 5: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("Added 1 to your rating: ", numRating+1)
	}

}
