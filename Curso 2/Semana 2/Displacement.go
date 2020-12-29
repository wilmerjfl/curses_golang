package main

import (
		"fmt"
		"bufio"
		"os"
		"strconv"
)

func GenDisplaceFn(a,vo, so float64) func(float64) float64{
	fn := func(t float64) float64 {
		return (0.5)*(a)*(t)*(t) + vo*t + so
	}	

	return fn
}

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	slice := make([] float64, 4)
	fmt.Println("Formula for displacement: s = 1/2*(a)*(t)^2 + vo*t + so")

	for i:=0; i < 4; i++ {
		switch i {
		case 0:
			fmt.Println("Enter a value for So")
		case 1:
			fmt.Println("Enter a value for a")
		case 2:
			fmt.Println("Enter a value for Vo")
		case 3:
			fmt.Println("Enter a value for t")
		}
		if scanner.Scan() {
			aux := scanner.Text()
			slice[i], _ = strconv.ParseFloat(aux, 64)
		}
	}
	fn := GenDisplaceFn(slice[0], slice[1], slice[2])

	fmt.Println(fn(slice[3]))
}