package main

import "time"

// isValidTimeSlot checks if the provided time slot is valid that is 1 hour.
func isValidTimeSlot(startTime, endTime string) (bool, error) {
	start, err := time.Parse("15:04", startTime)
	if err != nil {
		return false, err
	}
	end, err := time.Parse("15:04", endTime)
	if err != nil {
		return false, err
	}
	return end.Sub(start) == time.Hour, nil
}
