package main

import (
	"fmt"

	"pqan.com/learn-go/interface/cat"
	"pqan.com/learn-go/interface/dog"
	"pqan.com/learn-go/interface/service"
)

func getName(an service.AnimalName) {
	fmt.Printf("Cat name: %s \n", an.GetName())
}

func makeSound(a service.Animal) {
	fmt.Printf("%s keu sao? %s \n", a.GetName(), a.Speak())
}

func main() {
	// myDog := dog.Dog{
	// 	Name: "Noo",
	// }

	myDog, err := dog.New("Noo")
	if err != nil {
		panic(err)
	}
	getName(myDog)
	makeSound(myDog)

	// myCat := cat.Cat{
	// 	Name: "Bơ",
	// }

	myCat, err := cat.New("Bơ")
	if err != nil {
		panic(err)
	}
	getName(myCat)
	makeSound(myCat)
}
