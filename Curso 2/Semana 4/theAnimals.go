package main

import ("fmt"
				"os"
				"bufio"
			s "strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type newanimal struct {
	name string
	class string
}

type cow struct {
	food, locomotion, sound string
}
type bird struct {
	food, locomotion, sound string
}
type snake struct {
	food, locomotion, sound string
}

func (n cow) Eat() {
	println(n.food)
}
func (n cow) Move() {
	println(n.locomotion)
}
func (n cow) Speak() {
	println(n.sound)
}
func (n bird) Eat() {
	println(n.food)
}
func (n bird) Move() {
	println(n.locomotion)
}
func (n bird) Speak() {
	println(n.sound)
}
func (n snake) Eat() {
	println(n.food)
}
func (n snake) Move() {
	println(n.locomotion)
}
func (n snake) Speak() {
	println(n.sound)
}

func main() {
	slice := make([] newanimal, 0)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a command, must be:")
	fmt.Println("newanimal animal type")
	fmt.Println("Or")
	fmt.Println("query animal action")
	Cow := cow{"grass","walk","moo"}
	Bird := bird{"worms","fly","peep"}
	Snake := snake{"mice","slither","hsss"}
	for {
		fmt.Printf("> ")
		scanner.Scan()
		aux := scanner.Text()
		if (len(aux) > 0){
			animal := s.Split(aux," ")
			command, name, comm := animal[0], animal[1], animal[2]
			if (command == "newanimal"){
				slice = append(slice, NewAnimal(name, comm))
			} else if (command == "query"){

				for _, e := range(slice){
					if (e.name) == name {

						switch e.class {
						case "cow":
							switch comm{
							case "eat":
								Cow.Eat()
							case "move":
								Cow.Move()
							case "speak":
								Cow.Speak()
							}
						case "bird":
							switch comm{
							case "eat":
								Bird.Eat()
							case "move":
								Bird.Move()
							case "speak":
								Bird.Speak()
							}
						case "snake":
							switch comm{
							case "eat":
								Snake.Eat()
							case "move":
								Snake.Move()
							case "speak":
								Snake.Speak()
							}
						}
					}
					
				}
			}
		} else {
			break
		}
	
	}
}

func NewAnimal(name, class string) newanimal {
	a := newanimal{name, class}
	return a
}