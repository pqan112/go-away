package main

import (
	"cmp"
	"fmt"
)

func Max[T cmp.Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(1, 2))
	fmt.Println(Max(1, 1))
}
