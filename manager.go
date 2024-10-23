package main

import (
	"errors"
	"fmt"
)

// BookingManager interface defines methods for managing bookings.
type BookingManager interface {
	RegisterShow(name, genre string)
	OnboardShowSlots(name, slots string) error
	ShowAvailableByGenre(genre string) []string
	BookTicket(user, showName, timeSlot string, seats int) (int, error)
	CancelBooking(user string) error
	TrendingShow() (string, int)
}

type BookingSystem struct {
	shows         map[string]*Show
	nextBookingID int
}

func NewBookingSystem() *BookingSystem {
	return &BookingSystem{
		shows:         make(map[string]*Show),
		nextBookingID: 1,
	}
}

// ShowAvailableByGenre returns available shows for a given genre.
func (bs *BookingSystem) ShowAvailableByGenre(genre string) []string {

	var availableShows []string
	for _, show := range bs.shows {
		if show.Genre == genre {
			for timeSlot, capacity := range show.TimeSlots {
				if capacity > 0 {
					availableShows = append(availableShows, fmt.Sprintf("%s: (%s) %d", show.Name, timeSlot, capacity))
				}
			}
		}
	}
	return availableShows
}

// BookTicket allows a user to book tickets for a show.
func (bs *BookingSystem) BookTicket(user, showName, timeSlot string, seats int) (int, error) {

	show, exists := bs.shows[showName]
	if !exists {
		return 0, errors.New("show not found")
	}

	capacity, available := show.TimeSlots[timeSlot]
	if !available {
		return 0, errors.New("time slot not available")
	}

	if capacity < seats {
		// Add user to waitlist if not enough capacity
		show.Waitlist = append(show.Waitlist, user)
		bookingID := bs.nextBookingID
		bs.nextBookingID++
		show.Bookings[user] = Booking{
			ID:       bookingID,
			User:     user,
			TimeSlot: timeSlot,
			Seats:    seats,
			Status:   Waitlisted,
		}
		fmt.Printf("Booking Id : %d is in Waitlisting\n", bookingID)
		return bookingID, nil
	}

	// Book the tickets
	show.TimeSlots[timeSlot] -= seats
	booking := Booking{
		ID:       bs.nextBookingID,
		User:     user,
		TimeSlot: timeSlot,
		Seats:    seats,
		Status:   Confirmed,
	}
	show.Bookings[user] = booking
	show.BookedCount += seats
	bs.nextBookingID++

	fmt.Printf("Ticket Booked, Booking id: %d\n", booking.ID)
	return booking.ID, nil
}

// CancelBooking cancels a user's booking.
func (bs *BookingSystem) CancelBooking(user string) error {
	for showName, show := range bs.shows {
		if booking, exists := show.Bookings[user]; exists {
			// Free up the booked seats
			_ = show.TimeSlots[booking.TimeSlot]
			show.TimeSlots[booking.TimeSlot] += booking.Seats
			delete(show.Bookings, user)
			show.BookedCount -= booking.Seats
			booking.Status = Canceled
			fmt.Println("Booking Canceled")

			// Handle waitlist if available
			if len(show.Waitlist) > 0 {
				nextUser := show.Waitlist[0]
				show.Waitlist = show.Waitlist[1:] // Remove the user from waitlist
				// Try to book for the next user in waitlist
				if _, err := bs.BookTicket(nextUser, showName, booking.TimeSlot, booking.Seats); err == nil {
					fmt.Printf("Waitlist: %s has been booked for %s\n", nextUser, showName)
				}
			}
			return nil
		}
	}

	return errors.New("no booking found for user")
}

// TrendingShow returns the show with the most tickets booked.
func (bs *BookingSystem) TrendingShow() (string, int) {
	var trendingShow *Show
	for _, show := range bs.shows {
		if trendingShow == nil || show.BookedCount > trendingShow.BookedCount {
			trendingShow = show
		}
	}
	if trendingShow != nil {
		return trendingShow.Name, trendingShow.BookedCount
	}
	return "", 0
}
