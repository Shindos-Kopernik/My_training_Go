package main

import "fmt"

func main() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	someNames := allNames[1:3]
	allNames = append(allNames, "Gloves")
	allNames[1] = "Canoe"
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}
