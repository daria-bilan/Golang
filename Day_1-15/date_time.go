package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Getting current time
	//now := time.Now()
	//fmt.Println("Current time:", now)
	//fmt.Printf("Type: %T\n", now)
	//
	//// Access individual components
	//fmt.Println("\n=== Components ===")
	//fmt.Println("Year:", now.Year())
	//fmt.Println("Month:", now.Month())
	//fmt.Println("Month (int):", int(now.Month()))
	//fmt.Println("Day:", now.Day())
	//fmt.Println("Hour:", now.Hour())
	//fmt.Println("Minute:", now.Minute())
	//fmt.Println("Second:", now.Second())
	//fmt.Println("Nanosecond:", now.Nanosecond())
	//fmt.Println("Weekday:", now.Weekday())
	//fmt.Println("Day of year:", now.YearDay())

	//// Get data and time separately
	//year, month, day := now.Date()
	//fmt.Printf("\nDate: %d-%d-%d\n", year, month, day)
	//
	//hour, minute, second := now.Clock()
	//fmt.Printf("Time: %02d:%02d:%02d\n", hour, minute, second)

	// 2. Formatting Time
	// 🎯 IMPORTANT: Go uses reference time for formatting
	// Reference: Mon Jan 2 15:04:05 MST 2006
	// Think: 01/02 03:04:05PM '06 -0700

	//fmt.Println("=== Common Formats ===")
	//fmt.Println("YYYY-MM_DD:", now.Format("2006-01-02"))
	//fmt.Println("DD/MM/YYYY:", now.Format("02/01/2006"))
	//fmt.Println("MM/DD/YYYY:", now.Format("01/02/2006"))
	//fmt.Println("Full date:", now.Format("Monday, January 2, 2006"))
	//fmt.Println("Time 24h:", now.Format("15:04:05"))
	//fmt.Println("Time 12h:", now.Format("3:04:05 PM"))
	//fmt.Println("DateTime:", now.Format("2006-01-02 15:04:05"))
	//fmt.Println("ISO 8601:", now.Format("2006-01-02T15:04:05-0700"))
	//
	//// Custom formats
	//fmt.Println("\n=== Custom formats ===")
	//fmt.Println(now.Format("Jan 2, 2006 at 3:04PM"))
	//fmt.Println(now.Format("02.01.2006 в15:04"))
	//fmt.Println(now.Format("Monday, 02-Jan-06"))
	//
	//// Predefined constants
	//fmt.Println("\n=== Predefined constants ===")
	//fmt.Println("RFC3339:", now.Format(time.RFC3339))
	//fmt.Println("RFC822:", now.Format(time.RFC822))
	//fmt.Println("Kitchen:", now.Format(time.Kitchen))

	// 3. Parsing time
	// time.Parse(layout, value)
	// Layout must match the format of the value string

	//// Example 1. Standard date format
	//dateStr1 := "2025-10-26"
	//parsed1, err1 := time.Parse("2006-01-02", dateStr1)
	//if err1 != nil {
	//	fmt.Println("Error:", err1)
	//} else {
	//	fmt.Println("Parsed date:", parsed1)
	//	fmt.Printf("Type of dataStr: %T\n", dateStr1)
	//	fmt.Printf("Type of parsed1: %T\n", parsed1)
	//}

	//// Example 2. Date with time
	//dateStr2 := "2025-10-26 15:30:45"
	//parsed2, err2 := time.Parse("2006-01-02 15:04:05", dateStr2)
	//if err2 != nil {
	//	fmt.Println("Error:", err2)
	//} else {
	//	fmt.Println("Parsed datatime:", parsed2)
	//}

	//// Example 3. Different format
	//dateStr3 := "26/10/2025"
	//parsed3, err3 := time.Parse("02/01/2006", dateStr3)
	//if err3 != nil {
	//	fmt.Println("Error:", err3)
	//} else {
	//	fmt.Println("Parsed (DD/MM/YYYY):", parsed3)
	//}

	//// Example 4. With month name
	//dateStr4 := "Oct 26, 2025"
	//parsed4, err4 := time.Parse("Jan 2, 2006", dateStr4)
	//if err4 != nil {
	//	fmt.Println("Error:", err4)
	//} else {
	//	fmt.Println("Parsed (month name):", parsed4)
	//}

	//// Example 5. 12-hour format with AM/PM
	//dateStr5 := "3:30:45 PM"
	//parsed5, err5 := time.Parse("3:04:05 PM", dateStr5)
	//if err5 != nil {
	//	fmt.Println("Error:", err5)
	//} else {
	//	fmt.Println("Parsed time:", parsed5)
	//}

	//// Example 6. Parse with timezone
	//dateStr6 := "2025-10-26T15:30:45+02:00"
	//parsed6, err6 := time.Parse(time.RFC3339, dateStr6)
	//if err6 != nil {
	//	fmt.Println("Error:", err6)
	//} else {
	//	fmt.Println("Parsed ISO 8601:", parsed6)
	//}

	//// 4. Creating Specific Times
	//// Example 1: New Year 2025
	//newYear := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	//fmt.Println("New Year 2025:", newYear)
	//
	//// Example 2: My birthday
	//birthday := time.Date(1990, time.May, 15, 12, 30, 0, 0, time.UTC)
	//fmt.Println("Birthday:", birthday)
	//
	//// Example 3: Using local timezone
	//local := time.Date(2025, 10, 26, 15, 30, 0, 0, time.UTC)
	//fmt.Println("Local time:", local)
	//
	//// Example 4: Month as Integer
	//date := time.Date(2025, 10, 25, 0, 0, 0, 0, time.UTC)
	//fmt.Println("Date:", date)

	// 5. Time arithmetic (Duration)
	now := time.Now()
	fmt.Println("Now:", now.Format("2006-01-02 15:04:05"))

	//// Add duration
	//oneHourLater := now.Add(time.Hour)
	//fmt.Println("1 hour later:", oneHourLater.Format("2006-01-02 15:04:05"))
	//
	//thirtyMinLater := now.Add(30 * time.Minute)
	//fmt.Println("30 minutes later:", thirtyMinLater.Format("2006-01-02 15:04:05"))
	//
	//tenSecLater := now.Add(10 * time.Second)
	//fmt.Println("10 sec later:", tenSecLater.Format("2006-01-02 15:04:05"))

	//// Subtract duration (use negative)
	//oneHourAgo := now.Add(-time.Hour)
	//fmt.Println("1 hour ago:", oneHourAgo.Format("2006-01-02 15:04:05"))
	//
	//// Add days (24 hours)
	//tomorrow := now.Add(time.Hour * 24)
	//fmt.Println("Tomorrow:", tomorrow.Format("2006-01-02 15:04:05"))
	//
	//// Add date components (use AddDate for years, months, days)
	//nextWeek := now.AddDate(0, 0, 7)
	//fmt.Println("Next week:", nextWeek.Format("2006-01-02 15:04:05"))
	//
	//nextMonth := now.AddDate(0, 1, 0)
	//fmt.Println("Next month:", nextMonth.Format("2006-01-02 15:04:05"))
	//
	//nextYear := now.AddDate(1, 0, 0)
	//fmt.Println("Next year:", nextYear.Format("2006-01-02 15:04:05"))
	//
	//previousYear := now.AddDate(-1, 0, 0)
	//fmt.Println("Previous year:", previousYear.Format("2006-01-02 15:04:05"))
	//
	//// Combining
	//future := now.AddDate(1, 2, 15).Add(3 * time.Hour)
	//fmt.Println("1y 2m 15d 3h later:", future.Format("2006-01-02 15:04"))

	// 6. Comparing Times
}
