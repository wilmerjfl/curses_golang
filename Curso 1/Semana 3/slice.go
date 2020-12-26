package main

import (
	"fmt"
	s "strings"
	"strconv"
	"sort"
	)

func main() {
	
	Sli := make([]int, 3)
	Aux := len(Sli)
	for  i :=0; i < len(Sli) + 1; i++ {
		fmt.Println("Enter a integer number ")
		var Str string
		fmt.Scan(&Str)
		Num, err := strconv.Atoi(Str)
		if err == nil {
			//fmt.Println(i, len(Sli))
			if Aux > i {
				Sli[0] = Num
			} else {
				Sli = append(Sli, Num)
			}
		}
		sort.Ints(Sli)
		fmt.Println("-------------------------------------------")
		fmt.Println(Sli)
		if  s.ToLower(Str) == "x" {
			break
		}
	}
}