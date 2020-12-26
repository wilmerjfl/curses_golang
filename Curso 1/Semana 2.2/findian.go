package main

import ("fmt"
			s "strings"
			"bufio"
			"os"
)

func main() {

	fmt.Println("Enter a string")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {

		String := scanner.Text()
		Str := s.ToLower(s.ReplaceAll(String, " ",""))
		Contains := s.Contains(Str, "a")
		HasPrefix := s.HasPrefix(Str, "i")
		HasSuffix := s.HasSuffix(Str, "n")

		if HasPrefix && Contains && HasSuffix{
				fmt.Println("Found!")
		}else {
			fmt.Println("Not Found!")
		}
	}
}