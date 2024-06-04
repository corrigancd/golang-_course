package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct { // use capitalisation here because we don't want to expoe it to outside of the package
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     // this is inheritance in Go!! Currently it's assined as anon user
}

func (u User) OutputUserDetails() {
	fmt.Println("Printing user details based (function attached to struct) ...")
	// we don't need deferencing for structs, see pointers.go in section 4. This is a short cut for struct types
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() { // passed ad pointer so we can clear the user name from the user struct outside of this function
	fmt.Println("Clearing user first and last name...")
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("undefined first name, last name or birthdate")
	}

	// you can return a pointer here, example:
	// return &user{
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
