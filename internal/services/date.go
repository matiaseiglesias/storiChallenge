package services

import (
	"fmt"
	"time"
)

// GetMonthName extracts the month name from a date string.
// The function returns the full month name based on the provided date string.
//
// Parameters:
//   - date: A string representing the date in "dd/MM" format.
//
// Returns:
//   - string: The full name of the month extracted from the date. If the date is invalid,
//     an empty string is returned.
//
// Example:
//   month := GetMonthName("15/08") // Returns "August"
//
// Error Handling:
//   - If the date string cannot be parsed, an error message is printed, and the function
//     returns an empty string.

func GetMonthName(date string) string {
	// Define the layout for parsing the date
	layout := "02/01"

	// Parse the date string
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	// Get the full month name
	monthName := parsedDate.Format("January")

	return monthName
}
