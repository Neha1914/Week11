package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Addition:", Add(5, 3))
	fmt.Println("Subtract:", Subtract(6, 3))
	fmt.Println("Multiplication:", Multiply(5, 10))
}
