package main

import "fmt"

func main() {

	final := max(50, 20)

	fmt.Println("number1 is maximum ", final)

}

func max(num1, num2 int) bool {

	result := false

	if num1 > num2 {
		result = true
		return result

	} else {
		result = false
		return result

	}
	return result

}
