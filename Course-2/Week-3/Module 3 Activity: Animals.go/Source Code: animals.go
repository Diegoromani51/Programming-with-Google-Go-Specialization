package main

import ("fmt"
		"os")

type Animal struct {
	food string
	locomotion string
	sound string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	data := map[string]Animal{
		"bird": Animal{"worms", "fly", "peep"},
		"cow": Animal{"grass", "walk", "moo"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	for {
		fmt.Print(">")
		subjectAnimal := "0"
		animalAction := "0"
		_, err := fmt.Scan(&subjectAnimal, &animalAction)
		if err != nil {
			fmt.Println("Invalid input:", err)
			os.Exit(0)
		}
		if animalAction =="eat" {
			data[subjectAnimal].Eat()
		} else if animalAction == "move"{
			data[subjectAnimal].Move()
		} else if animalAction == "speak" {
			data[subjectAnimal].Speak()
		}

	}
}
