package main

import "fmt"

func main() {
	fmt.Printf("Enter a Float Number To convert in a Int \n")
	var Num float64
	fmt.Scan(&Num)
	fmt.Printf("The number is: %d",int(Num))
}