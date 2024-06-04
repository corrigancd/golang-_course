package main

import (
	"fmt"
	"time"
)

type user struct { // use lowercase here because we don't want to expoe it to outside of the package
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println("Printing user details based (function attached to struct) ...")
	// we don't need deferencing for structs, see pointers.go in section 4. This is a short cut for struct types
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserName() { // passed ad pointer so we can clear the user name from the user struct outside of this function
	fmt.Println("Clearing user name...")
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) user {
	// you can return a pointer here, example:
	// return &user{
	return user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// regular notation:

	var appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
	outputUserDetailsSeparateFunction(&appUser) // logging as an independent function
	appUser.outputUserDetails()                 // logging when adding the logging function to the struct
	appUser.clearUserName()
	appUser.outputUserDetails()

	// if the order is in the same as the blueprint, you can use the shorthand notation below:
	var appUserShortNotation user
	var createdAt = time.Now()
	appUserShortNotation = user{
		firstName,
		lastName,
		birthdate,
		createdAt,
	}
	outputUserDetailsSeparateFunction(&appUserShortNotation) // logging as an independent function
	appUserShortNotation.outputUserDetails()                 // logging when adding the logging function to the struct

	var appUsersFromConstructor user
	appUsersFromConstructor = newUser(firstName, lastName, birthdate)
	fmt.Println("From Constructor!...")
	appUsersFromConstructor.outputUserDetails()

	// ... do something awesome with that gathered data!

}

func outputUserDetailsSeparateFunction(u *user) {
	fmt.Println("Printing user details (function not in struct) ...")
	// we don't need deferencing for structs, see pointers.go in section 4. This is a short cut for struct types
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
