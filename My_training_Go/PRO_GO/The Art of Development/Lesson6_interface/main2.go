package main

var _ User = superUser{}

type superUser struct {
	Name string
	Age  int
}

func (s *superUser) ChangeFIO(newFIO string) {
	//TODO implement me
	panic("implement me")
}

func (s *superUser) ChangeAddress(newAddress string) {
	//TODO implement me
	panic("implement me")
}

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
}

func (u *user) ChangeFIO(newFIO string) {
	u.FIO = newFIO
}

func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress
}

type User interface {
	ChangeFIO(newFIO string)
	ChangeAddress(newAddress string)
}

func NewUser(fio, address, phone string) {
	u := user{
		FIO:     fio,
		Address: address,
		Phone:   phone,
	}
}
func main() {

}
