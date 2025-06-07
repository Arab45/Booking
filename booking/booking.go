package booking

import "time"

type BookingUser struct {
	userId  string
	product make(map[string]string);
	date    time.Time
	delivery string
}

func Booking() {

	var collections = map[string]BookingUser{
		userId: userId,
		product: product,
		date: date,
		delivery: delivery
	}


}