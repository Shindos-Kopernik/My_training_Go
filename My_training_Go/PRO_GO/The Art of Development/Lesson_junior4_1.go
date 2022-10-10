package main

import "fmt"
import "github.com/mitchellh/mapstructure"

type Point struct {
	X int
	Y int
}

func (p Point) method() {
	fmt.Println("Call Point Metod")
}
func main() {
	pointsMap := map[string]int{
		"x": 1,
		"y": 2,
	}

}
