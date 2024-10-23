package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Show represents a show
type Show struct {
	Name        string
	Genre       string
	TimeSlots   map[string]int // TimeSlot -> Capacity
	Bookings    map[string]Booking
	Waitlist    []string
	BookedCount int
}

// RegisterShow registers a new show with name and genre.
func (bs *BookingSystem) RegisterShow(name, genre string) {
	bs.shows[name] = &Show{
		Name:      name,
		Genre:     genre,
		TimeSlots: make(map[string]int),
		Bookings:  make(map[string]Booking),
	}
	fmt.Printf("%s show is registered !!\n", name)
}

// OnboardShowSlots onboard the available time slots for a show.
func (bs *BookingSystem) OnboardShowSlots(name string, slots string) error {

	show, exists := bs.shows[name]
	if !exists {
		return errors.New("show not found")
	}

	slotEntries := strings.Split(strings.TrimSpace(slots), ",")
	for _, entry := range slotEntries {
		parts := strings.Fields(entry)
		if len(parts) != 2 {
			return errors.New("invalid slot format")
		}

		timeRange := parts[0]
		capacity, err := strconv.Atoi(parts[1])
		if err != nil || capacity <= 0 {
			return errors.New("invalid capacity")
		}

		slotParts := strings.Split(timeRange, "-")
		if len(slotParts) != 2 {
			return errors.New("invalid time range format")
		}

		startTime := slotParts[0]
		endTime := slotParts[1]

		valid, err := isValidTimeSlot(startTime, endTime)
		if err != nil || !valid {
			return errors.New("time slot must be exactly one hour")
		}

		timeSlot := fmt.Sprintf("%s-%s", startTime, endTime)
		if _, exists := show.TimeSlots[timeSlot]; exists {
			return errors.New("time slot already exists")
		}

		show.TimeSlots[timeSlot] = capacity
	}
	fmt.Println("Slots onboarded successfully!")
	return nil
}
