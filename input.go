package main

import (
	"fmt"
)

func main() {
	var fname string
	fmt.Println("Enter your Name :")
	fmt.Scan(&fname)

	fmt.Println("Your Name :" + fname)
}
