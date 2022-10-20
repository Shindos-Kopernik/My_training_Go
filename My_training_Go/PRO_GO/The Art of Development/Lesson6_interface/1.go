package main

import "fmt"

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type InterfaceHere interface {
	Sum() int
}

func main() {
	var a InterfaceHere
	sh := structHere{N1: 1, N2: 2}
	os := otherStruct{A: 2, B: 3}
	a = &sh
	a.Sum()
	fmt.Println(a.Sum())
	a = &os
	fmt.Println(a.Sum())
	var i InterfaceHere = otherStruct{A: 10, B: 43}
	fmt.Println(i.Sum())
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}
