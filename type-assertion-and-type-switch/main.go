package main

import "fmt"

// type assertion
func printStr(val interface{}) {
	str, ok := val.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("Hong phai kieu string roi")
	}
}

// type switch
func printValue(val any) {
	switch val.(type) {
	case int, string:
		fmt.Println(val)
	default:
		fmt.Println("Type invalid")
	}

}

func main() {
	printStr("Hehe")
	printStr(10)

	printValue("Hehe")
	printValue(10)
}
