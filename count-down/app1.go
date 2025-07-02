package main

import "fmt"

func countDown(number int) {
	fmt.Println(number)

	if number > 0 {
		countDown(number - 1)
	}
}

func main() {
	countDown(10)
}
