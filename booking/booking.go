package booking

import (
	"fmt"
	"time"
)

type BookingUser struct {
	userId  string
	product map[string]string;
	date    time.Time
	delivery string
}

func Booking() {
	var bookings = make([]BookingUser, 0)


	
	var userId string
	var product map[string]string
	var date time.Time
	var delivery string

	fmt.Printf("provide your userId:");
	fmt.Scan(&userId)
	fmt.Printf("provide your product order:");
	fmt.Scan(&product)
	fmt.Printf("current time:");
	fmt.Scan(&date)
	fmt.Printf("delivery day:");
	fmt.Scan(&delivery)



	var bookingTable = BookingUser{
		userId:  userId,
	    product: product,
	    date: date,
	    delivery: delivery,	
	}

	var collection = append(bookings, bookingTable);
	fmt.Printf("booking collection details %v\n", collection);


}