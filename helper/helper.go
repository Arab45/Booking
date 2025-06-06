package helper
import "strings"

func ValidateUserInput(firstName string, lastName string, email string, tickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := tickets > 0 && tickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}