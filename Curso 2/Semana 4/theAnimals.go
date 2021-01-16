package main

import ("fmt"
				"os"
				"bufio"
			s "strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct {
	name string
}
type bird struct {
	name string
}
type snake struct {
	name string
}
func (n *cow) Name(name string) {
	n.name = name
}
func (n *bird) Name(name string) {
	n.name = name
}
func (n *snake) Name(name string) {
	n.name = name
}
func (n cow) Eat() {
	println("grass")
}
func (n cow) Move() {
	println("walk")
}
func (n cow) Speak() {
	println("moo")
}
func (n bird) Eat() {
	println("worms")
}
func (n bird) Move() {
	println("fly")
}
func (n bird) Speak() {
	println("peep")
}
func (n snake) Eat() {
	println("mice")
}
func (n snake) Move() {
	println("slither")
}
func (n snake) Speak() {
	println("hsss")
}

func main() {
	
	animals := make(map[string] Animal)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a command, must be:")
	fmt.Println("newanimal animal type")
	fmt.Println("Or")
	fmt.Println("query animal action")

	for {
		fmt.Printf("> ")
		scanner.Scan()
		aux := scanner.Text()
		if (len(aux) > 0){
			animal := s.Split(aux," ")
			command, name, comm := animal[0], animal[1], animal[2]
			if (command == "newanimal"){
				switch comm {
				case "cow":
					animals[name] = cow{}
				case "bird":
					animals[name] = bird{}
				case "snake":
					animals[name] = snake{}
				}
				
			} else if (command == "query"){

				switch comm{
				case "eat":
					animals[name].Eat()
				case "move":
					animals[name].Move()
				case "speak":
					animals[name].Speak()
				}
			}
		} else {
			break
		}
	
	}
}