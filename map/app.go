package main

import "fmt"

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	students := map[int]string{
		1: "An",
		2: "Minh",
		3: "Hoang",
	}

	fmt.Println(students[1])

	for k, v := range students {
		fmt.Printf("students[%d]: %s \n", k, v)
	}

	employees := map[string]Employee{
		"an":    {Name: "An", Age: 25, Role: "Dev quen"},
		"anime": {Name: "Anime", Age: 25, Role: "CEO"},
	}

	fmt.Println(employees["anime"].Role)
	for k, v := range employees {
		fmt.Println(k)
		fmt.Println(v.Age)
	}

	scores := map[string][]int{
		"math":    {90, 80, 85},
		"english": {75, 88},
	}
	fmt.Println(scores["math"][0])
	scores["math"] = append(scores["math"], 95)

	for _, v := range scores {
		for _, score := range v {
			fmt.Println(score)
		}
	}

}
