package services

import (
	"fmt"
	"time"
)

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

	// Output the result
	fmt.Println("Date:", date)
	fmt.Println("Month Name:", monthName)
	return monthName
}
