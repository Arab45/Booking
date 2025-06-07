package booking

import (
	"fmt"
	"time"
)

type BookingUser struct {
	UserId   string
	Product  map[string]string
	Date     time.Time
	Delivery string
}

func Booking() {
	var bookings = make([]BookingUser, 0)

	var userId string
	var delivery string

	// --- For now, we'll simulate input for map and date ---
	// because you cannot scan a map or time.Time directly from the console

	// Simple input for userId
	fmt.Printf("Provide your userId: ")
	fmt.Scan(&userId)

	// Simulated product map input
	product := make(map[string]string)
	var productKey, productValue string

	fmt.Printf("Enter product key: ")
	fmt.Scan(&productKey)
	fmt.Printf("Enter product value: ")
	fmt.Scan(&productValue)
	product[productKey] = productValue

	// Set current time for booking
	date := time.Now()

	// Delivery input
	fmt.Printf("Delivery day: ")
	fmt.Scan(&delivery)

	// Construct the booking
	var bookingTable = BookingUser{
		UserId:   userId,
		Product:  product,
		Date:     date,
		Delivery: delivery,
	}

	// Append to collection
	bookings = append(bookings, bookingTable)

	// Print the result
	fmt.Printf("Booking collection details:\n%+v\n", bookings)
}
