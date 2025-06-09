package profile

import "fmt"

func Profile() {
	type Profile struct {
		firstName string
		lastName  string
		userName  string
		email     string
		phone     int
		address   string
		password  string
	}

	var collection []Profile

	var firstName string
	var lastName string
	var userName string
	var email string
	var phone int
	var address string
	var password string

	fmt.Print("Enter your first name:");
	fmt.Scan(&firstName);
	fmt.Print("Enter your last name:");
	fmt.Scan(&lastName);
	fmt.Print("Enter your email address:");
	fmt.Scan(&email);
	fmt.Print("Enter your user name:");
	fmt.Scan(&userName);
	fmt.Print("Enter your phone number:");
	fmt.Scan(&phone);
	fmt.Print("Enter your address:");
	fmt.Scan(&address);
	fmt.Print("Enter your password:");
	fmt.Scan(&password);

	var profileTable = Profile{
		firstName: firstName,
		lastName:  lastName,
		userName:  userName,
		email:     email,
		phone:     phone,
		address:   address,
		password:  password,
	}
}