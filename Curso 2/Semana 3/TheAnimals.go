package main

import ("fmt")



type animal struct {
	food string
	locomotion string
	sound string
}

func main() {
	var aux string
	var att string
	var pick animal

	var cow animal
	cow = New("grass", "walk", "moo")
	var bird animal
	bird = New("worms","fly","peep")
	var snake animal
	snake = New("mice","slither","hsss")

	exit := true

	for (exit) {
		fmt.Println("Choose One Animal or Press 0 to exit")
		fmt.Println(">cow")
		fmt.Println(">bird")
		fmt.Println(">snake")
		fmt.Scan(&aux)
		switch aux {
		case "cow":
			pick = cow
		case "bird":
			pick = bird
		case "snake":
			pick = snake
		case "0":
			exit = false
			continue
		default:
			fmt.Println("Enter a validate value")
			continue
		}
		fmt.Println("Pick some attributte")
		fmt.Println(">eat")
		fmt.Println(">move")
		fmt.Println(">speak")
		fmt.Scan(&att)
		switch att {
		case "eat":
			pick.Eat()
		case "move":
			pick.Move()
		case "speak":
			pick.Speak()
		default:
			continue
		}
	}
}

func (a animal) Eat() {
	fmt.Println(a.food)
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
}

func (a animal) Speak() {
	fmt.Println(a.sound)
}

func New(food, move, speak string) animal {
	a := animal{food, move, speak}
	return a
}