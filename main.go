package main

import (
	"fmt"
)

// Main driver function to test the booking system.
func main() {
	bs := NewBookingSystem()

	// Example interactions

	// Register shows
	bs.RegisterShow("TMKOC", "Comedy")
	bs.RegisterShow("The Sonu Nigam Live Event", "Singing")
	bs.RegisterShow("Tech Innovations", "Tech")

	// Onboard show slots
	err := bs.OnboardShowSlots("Indias got Talent", "9:00-11:00")
	if err != nil {
		fmt.Println(err)
	}
	err = bs.OnboardShowSlots("TMKOC", "9:00-10:00 3, 12:00-13:00 2, 15:00-16:00 5")
	if err != nil {
		fmt.Println(err)
	}
	bs.OnboardShowSlots("The Sonu Nigam Live Event", "9:00-10:00 3, 13:00-14:00 2, 17:00-18:00 1")
	bs.OnboardShowSlots("Tech Innovations", "11:00-12:00 4, 14:00-15:00 5")

	// Show available by genre
	comedyShows := bs.ShowAvailableByGenre("Comedy")
	fmt.Println("Available Comedy Shows:")
	for _, show := range comedyShows {
		fmt.Println(show)
	}

	singingShows := bs.ShowAvailableByGenre("Singing")
	fmt.Println("Available Singing Shows:")
	for _, show := range singingShows {
		fmt.Println(show)
	}

	// Book tickets
	_, err = bs.BookTicket("UserA", "TMKOC", "12:00-13:00", 2)
	if err != nil {
		fmt.Println(err)
	}
	_, err = bs.BookTicket("UserB", "TMKOC", "12:00-13:00", 1) // Waitlisting
	if err != nil {
		fmt.Println(err)
	}
	_, err = bs.BookTicket("UserB", "TMKOC", "15:00-16:00", 1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = bs.BookTicket("UserB", "TMKOC", "15:00-16:00", 1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = bs.BookTicket("UserB", "TMKOC", "15:00-16:00", 4) // Waitlisting
	if err != nil {
		fmt.Println(err)
	}
	_, err = bs.BookTicket("UserB", "TMKOC", "15:00-16:00", 3)
	if err != nil {
		fmt.Println(err)
	}
	_, err = bs.BookTicket("UserB", "TMKOC", "15:00-16:00", 1)
	if err != nil {
		fmt.Println(err)
	}
	name, count := bs.TrendingShow()
	fmt.Printf("Trending Show: %s with %d tickets booked.\n", name, count)

	// Show available after booking
	comedyShows = bs.ShowAvailableByGenre("Comedy")
	fmt.Println("Available Comedy Shows after booking:")
	for _, show := range comedyShows {
		fmt.Println(show)
	}

	// Cancel booking
	_ = bs.CancelBooking("UserA")

	// Show available after cancellation
	comedyShows = bs.ShowAvailableByGenre("Comedy")
	fmt.Println("Available Comedy Shows after cancellation:")
	for _, show := range comedyShows {
		name, count = bs.TrendingShow()
		fmt.Printf("Trending Show: %s with %d tickets booked.\n", name, count)
		fmt.Println(show)
	}

	// Book another ticket from waitlist
	_, _ = bs.BookTicket("UserB", "TMKOC", "12:00", 1)
}
