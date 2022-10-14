package main

import "fmt"

func printPrice() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
}
func main() {
	fmt.Println("About to call function")
	printPrice()
	fmt.Println("Function complete")
}
