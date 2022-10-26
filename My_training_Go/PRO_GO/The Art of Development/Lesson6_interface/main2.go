package main

var _ User = &superUser{}

type superUser struct {
	Name      string
	Age       int
	IsBlocked bool
}

func (s *superUser) Block() {
	u.IsBlocked = true
}

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
	IsBlocked           bool
}

func (u *user) Block() {
	u.IsBlocked = true
}

type User interface {
	Block()
}

func NewUser(fio, address, phone string) User {
	u := user{}
	return &u
}
func main() {
	u := NewUser("", "", "")
	u.Block()
}
