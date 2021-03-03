// Dates Range, https://tpkn.me
package datesrange

import (
	"errors"
	"fmt"
	"time"
)

const layout string = "2006-01-02"

func New(start, end string) ([]time.Time, error) {
	if start == "" && end == "" {
		return nil, errors.New("Invalid dates range. Make sure that the start and end dates are not empty at the same time!")
	}
	// Check if the dates are not empty
	if start == "" {
		return nil, errors.New("Invalid dates range. Make sure that the start date is not empty!")
	}
	
	// If 'end' is empty, get today's date
	if end == "" {
		end = time.Now().Format(layout)
	}
	
	// Convert string to YYYY-MM-DD format time
	start_time, _ := time.Parse(layout, start)
	end_time, _ := time.Parse(layout, end)
	
	// Check if the dates are correct
	if end_time.Sub(start_time) < 0 {
		return nil, errors.New(fmt.Sprintf("Invalid dates passed '%s:%s'", start, end))
	}
	
	result := []time.Time { start_time }
	
	if start_time == end_time {
		return result, nil
	}
	
	// Adding days until it hits the "end_date" date
	for start_time != end_time {
		start_time = start_time.AddDate(0, 0, 1)
		result = append(result, start_time)
	}
	
	return result, nil
}
