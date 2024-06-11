package main

import (
	"errors"
	"fmt"
	"time"

	"example.com/structs/user"
)

type userLocal struct { // use lowercase here because we don't want to expoe it to outside of the package
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u userLocal) outputUserDetails() {
	fmt.Println("Printing user details based (function attached to struct) ...")
	// we don't need deferencing for structs, see pointers.go in section 4. This is a short cut for struct types
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *userLocal) clearUserName() { // passed ad pointer so we can clear the user name from the user struct outside of this function
	fmt.Println("Clearing user name...")
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*userLocal, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("undefined first name, last name or birthdate")
	}
	// you can return a pointer here, example:
	// return &user{
	return &userLocal{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// regular notation:

	var appUser = userLocal{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
	OutputUserDetailsSeparateFunction(&appUser) // logging as an independent function
	appUser.outputUserDetails()                 // logging when adding the logging function to the struct
	appUser.clearUserName()
	appUser.outputUserDetails()

	// if the order is in the same as the blueprint, you can use the shorthand notation below:
	var appUserShortNotation userLocal
	var createdAt = time.Now()
	appUserShortNotation = userLocal{
		firstName,
		lastName,
		birthdate,
		createdAt,
	}
	OutputUserDetailsSeparateFunction(&appUserShortNotation) // logging as an independent function
	appUserShortNotation.outputUserDetails()                 // logging when adding the logging function to the struct

	// var appUsersFromConstructor *user
	var appUsersFromConstructor, err = newUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("From Constructor!...")
	appUsersFromConstructor.outputUserDetails()

	// Creating user from a package
	var appUsersFromConstructorInAnotherPackage, err2 = user.New(firstName, lastName, birthdate)
	if err2 != nil {
		fmt.Println("ERROR: ", err2)
		return
	}
	fmt.Println("From Constructor from another package!...")
	appUsersFromConstructorInAnotherPackage.OutputUserDetails()

	admin := user.NewAdmin("admin@example.com", "12345")
	fmt.Println("From Constructor from another package!...")
	admin.OutputUserDetails() // anon means that we can access the embedded struct directly from the parent struct instead

}

func OutputUserDetailsSeparateFunction(u *userLocal) {
	fmt.Println("Printing user details (function not in struct) ...")
	// we don't need deferencing for structs, see pointers.go in section 4. This is a short cut for struct types
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
