package main

import "fmt"

func main() {
	school := [][] string {
		{"A", "B", "C"},
		{"H", "N", "T"},
	}

	for _, class := range school {
		for _, student := range class {
			fmt.Println(student)
		}
	}

	apple1 := []string {"Apple 1", "Apple 2", "Apple 3"}
	apple2 := []string {"Apple 4", "Apple 5", "Apple 6"}
	apple3 := []string {"Apple 7", "Apple 8", "Apple 9"}
	//apple1 = append(apple1, apple2..., apple3...) ❌
	apple1 = append(apple1, apple2...)
	apple1 = append(apple1, apple3...)
	fmt.Print(apple1)
}

