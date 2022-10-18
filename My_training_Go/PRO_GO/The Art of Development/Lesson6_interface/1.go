package main

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() {
	return s.N1 + s.N2
}

type Interface interface {
	Sum() int
}

func main()  {

}
type otherStruct struct {
	A, B int()
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}