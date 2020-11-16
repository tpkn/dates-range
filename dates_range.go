//
// Dates Range, https://tpkn.me
//
package datesrange

import (
	"errors"
	"time"
)

const layout string = "2006-01-02"

func New(start, end string) ([]time.Time, error) {
	
	// Convert string to YYYY-MM-DD format time
	start_time, _ := time.Parse(layout, start)
	
	// If 'end' is empty, get today's date
	if end == "" {
		end = time.Now().Format(layout)
	}
	end_time, _ := time.Parse(layout, end)
	
	// Check if the dates are correct
	if end_time.Sub(start_time) < 0 {
		return nil, errors.New(fmt.Sprintf("Invalid dates passed '%s:%s'", start, end))
	}
	
	// Adding days until it hits the "end_date" date
	start_date := time.Date(start_time.Year(), start_time.Month(), start_time.Day(), 0,0,0,0, time.Local)
	end_date := time.Date(end_time.Year(), end_time.Month(), end_time.Day(), 0,0,0,0, time.Local)
	
	result := []time.Time { start_date }
	
	if start_date == end_date {
		return result, nil
	}
	
	for start_date != end_date {
		start_date = start_date.AddDate(0, 0, 1)
		result = append(result, start_date)
	}
	
	return result, nil
}
