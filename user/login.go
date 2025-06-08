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

	fmt.Printf("insert your username:")
	fmt.Scan(&username);
	fmt.Printf("password most be mixture of upperCase and lowerCase:")
	fmt.Scan(&password);

	var loginDetails = Login{
		username: username,
		password: password,
	}

	var userTable = append(collection, loginDetails);
	fmt.Printf("user login collection %v\n", userTable);
}