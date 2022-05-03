package cours

import "fmt"

type User struct {
	firstName string
}

func (u *User) Hello() {
	fmt.Println("Hello", u.firstName)
}

func MakeUser(firstName string) *User {
	return &User{
		firstName: firstName,
	}
}
