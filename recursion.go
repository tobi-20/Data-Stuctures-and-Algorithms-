package main

import "fmt"

func Factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * Factorial(x-1)
	}

}

func main() {
	fmt.Println(Factorial(4))
}
