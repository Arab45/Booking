package main
import "fmt"


func main() {

	var conferenceName string = "Go conference";
	const conferenceTickets int = 50;
	var remainingTickets int = 50;
	var bookings [50]string;
	var slices []string;

	slices = append(slices, "John Smith");
	fmt.Println("slice in Go is also an array", slices[0]);

	fmt.Println(&conferenceName);
	fmt.Println(&remainingTickets);

	fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName);

	fmt.Printf("Welcome to %v booking application\n", conferenceName);
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets);
	fmt.Println("Get your tickets here to attend");	

	var firstName string;
	var lastName string;
	var email string;
	var tickets int;

	fmt.Println("Enter your first name: ");
	fmt.Scan(&firstName);
fmt.Println("Enter your last name: ");
	fmt.Scan(&lastName);
	fmt.Println("Enter your email: ");
	fmt.Scan(&email);
	fmt.Println("Enter number of tickets: ");
	fmt.Scan(&tickets);
	fmt.Printf("User %v %v with email %v booked %v tickets\n", firstName, lastName, email, tickets);

	var userName string;
	var userTickets int;
	fmt.Println("Enter your name: ");
	fmt.Scan(&userName);
	fmt.Println("Enter number of tickets: ");
    fmt.Scan(&userTickets);
	fmt.Println(&userName);
	fmt.Println(&userTickets);
	if userTickets > 0{
		var leftTickets int = remainingTickets - userTickets;
		fmt.Printf("Booking tickets left to purchase: %v\n", leftTickets); 
	};

	bookings[0] = firstName + " " + lastName;
	fmt.Printf("List of bookings is %v\n", bookings);
	fmt.Printf("Type of bookings is %T\n", bookings);
	fmt.Printf("First booking is %v\n", bookings[0]);
	fmt.Printf("Number of bookings is %v\n", len(bookings));

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets);
}