package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ {
		i = i
		PrintNumber(i)
	}
}
func PrintHello() {
	fmt.Println("Hello, Go")
}

// PrintNumber writes a number using the fmt.Println function
func PrintNumber(number int) {
	fmt.Println(number)
}
