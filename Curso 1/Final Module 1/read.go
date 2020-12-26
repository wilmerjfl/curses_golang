package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main()  {
	var File string
	type name struct {
		fname string
		lname string
	}
	people := make([] name, 0)
	p := fmt.Println
	p("Enter name of file")
	fmt.Scan(&File)
	//p(File)
	c, err := os.Open(File+".txt")

	if err != nil {
		p("An error has been appear %s", err)
	}
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), " ")
		people = append(people, name{names[0], names[1]})
	}
	
	for _, v := range people {
		p(v.fname, v.lname)
	}
	//p(names)
	c.Close()
}