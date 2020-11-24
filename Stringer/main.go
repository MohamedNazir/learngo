package main

import "fmt"

func main() {

	p := person{"Mohamed", 20, "test@gmail.com"}

	fmt.Println(p.String())
}

type person struct {
	name  string
	age   int
	email string
}

func (u *person) String() string {
	return fmt.Sprintf("%s, %d, <%s>", u.name, u.age, u.email)
}
