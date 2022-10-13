package main

import "fmt"

func main() {
	product := map[string]float64{
		"Kayak": 1,
		"Lol":   40.12,
		"Hat":   0,
	}
	value, ok := product["Hat"]
	if ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No")
	}
	fmt.Println(product["Kayak"])
	fmt.Println(product["Lol"])
}
