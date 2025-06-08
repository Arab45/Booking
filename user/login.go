package login

import "fmt"

func Login() {
	type Login struct {
		username string
		password string
	}

	var collection = make([]Login, 0)

	var username string
	var password string

	fmt.Printf("")

	var loginDetails = Login{
		username: username,
		password: password,
	}
}