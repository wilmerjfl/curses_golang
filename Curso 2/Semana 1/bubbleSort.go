package main

import ("fmt"
		"bufio"
		"os"
		"strconv"
)

func main() {

	
	slice := make([]int,10)
	scanner := bufio.NewScanner(os.Stdin)
	for i :=0; i < 10; i++ {
		fmt.Println("Enter a Integer")
		if scanner.Scan() {
			integer := scanner.Text()
			slice[i], _ = strconv.Atoi(integer)
		}
	}
	BubbleSort(slice)
	fmt.Println(slice)
}

func BubbleSort(sli []int) int  {
	order := true
	for order {
		order = false
		for i := 1; i < len(sli); i++ {
            if sli[i-1] > sli[i] {
				Swap(sli, i)
                order = true
            }
        }
	}
	return 0
}

func Swap(s []int, i int) int {
	s[i], s[i-1] = s[i-1], s[i]
	return 0
}