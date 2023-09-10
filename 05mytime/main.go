package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my study!")
	presentTime := time.Now()
	fmt.Printf("%v\n", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}
