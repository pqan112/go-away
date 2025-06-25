package main

import "fmt"

func main() {
	/**
	exercise 1:
	- print odd number from 1 to 100
	- keep 3 number in a line
	- each number is separated by a comma
	*/

	count := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			if count > 0 {
				fmt.Print(",")
			}
			fmt.Printf("%d", i)
			count++
			if count == 3 {
				fmt.Print("\n")
				count = 0
			}
		}
	}

	if count > 0 {
		fmt.Print("\n")
	}

	/**
	exercise 2:
	- print multiplication table
	*/

	var start, end int

	fmt.Print("nhap so bat dau: ")
	fmt.Scanf("%d", &start)

	fmt.Print("nhap so ket thuc: ")
	fmt.Scanf("%d", &end)

	if start == 0 || end == 0 {
		fmt.Print("So bat dau va so ket thuc phai lon hon 0")
	} else if start > end {
		fmt.Print("So ket thuc phai lon hon hoac bang so bat dau")
	} else {
		for m := start; m <= end; m++ {
			fmt.Printf("bang cuu chuong %d \n", m)
			for n := 1; n <= 10; n++ {
				fmt.Printf("%d x %d = %d \n", m, n, m*n)
			}
		}
	}

}
