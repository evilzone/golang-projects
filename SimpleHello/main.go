package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("The number is : %d\n", add(4, 5))
}
