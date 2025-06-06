package premium

import (
	"fmt"
)

func PremiumUser() {
	type UserPremiumDetails struct {
		userId    string
		userName  string
		isPremium bool
	}

	var user = make([]UserPremiumDetails, 0)

	var userId    string
	var userName  string
	var isPremium bool

	fmt.Printf("input your userId: ");
	fmt.Scan(&userId);
	fmt.Printf("input your userName: ");
	fmt.Scan(&userName);
	fmt.Printf("input your isPremium: ");
	fmt.Scan(&isPremium);

	var userDetails = UserPremiumDetails{
		userId:    userId,
		userName:  userName,
		isPremium: isPremium,
	}



	var userTable = append(user, userDetails)

	fmt.Printf("collection of user premium registry %v\n", userTable);

}