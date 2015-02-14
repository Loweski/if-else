// if_else project if_else.go
package main

import (
	"fmt"
)

type Person struct {
	name     string
	position string
	course   string
}
type myPrint func(string)

func Personprinter(list string) myPrint {
	return func(l string) {
		fmt.Print(l + list)
	}
}
func main() {

	var person1 = Person{"Leng", "Student", "Csci 130"}
	if true {
		Personprinter(person1.name)
	} else {
		fmt.Println("Person is not listed")
	}
	fmt.Println("OMG")

	var answer = 42

	switch answer {
	case 42:
		fmt.Println("You are correct.")
	case 22:
		fmt.Println("YOU ARE WRONG.")
	case 0:
		fmt.Println("YOU ARE VERY WRONG.")
	}
}
