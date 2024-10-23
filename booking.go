package main

// BookingStatus gives status of a booking.
type BookingStatus string

type Role string

// Booking represents a booking made by a user.
type Booking struct {
	ID       int
	User     string
	TimeSlot string
	Seats    int
	Status   BookingStatus
}

// User represents user as organizer or consumer/customer.
type User struct {
	Name           string
	AssociatedRole Role
}
