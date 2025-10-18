package main

// import (
// 	"fmt"
// )

// func main() {
// //Method 1: Declare with zero values
// var numbers [5]int
// fmt.Println(numbers)

// //Method 2: Declare and initialize
// var fruits [3]string = [3]string{"apple", "banana", "cherry"}
// fmt.Println(fruits)

// //Method 3: Short declaration with inferred size
// colors := [...]string{"red", "green", "blue"}
// fmt.Println(colors)

// //Method 4: Initialize specific indices
// grades := [5]int{0: 85, 2: 90, 4: 95}
// fmt.Println(grades)

// temperatures := [7]float64{22.5, 23.1, 24.0, 25.5, 23.8, 22.0, 21.5}

// // Access element
// fmt.Println("Monday temperature:", temperatures[0])

// // Modify element
// temperatures[3] = 26.0
// fmt.Println("Updated Wednesday:", temperatures[3])

// index out of range
// fmt.Println(temperatures[10])

// score := [4]int{85, 90, 78, 92}

// length := len(score)
// fmt.Println("Number of scores:", length)

// Length is fixed - compile error:
// score[4] = 88 // index out of range [4] with length 4

// Itirating
// cities := [4]string{"Kyiv", "Lviv", "Odesa", "Kharkiv"}

// // Method 1: Traditional for loop
// fmt.Println("Method 1: Index-based loop")
// for i := 0; i < len(cities); i++ {
// 	fmt.Printf("City %d: %s\n", i, cities[i])
// }

// // Method 2: Range-based loop (most idiomatic)
// fmt.Println("\nMethod 2: Range loop with index and value")
// for index, city := range cities {
// 	fmt.Printf("City %d: %s\n", index, city)
// }

// // Method 3: Range loop - value only
// fmt.Println("\nMethod 3: Range loop with value only")
// for _, city := range cities {
// 	fmt.Println(city)
// }

// // Methid 4: Range loop - index only
// fmt.Println("\nMethod 4: Range loop with index only")
// for index := range cities {
// 	fmt.Printf("Index: %d\n", index)
// }

// original := [3]int{1, 2, 3}
// copy := original

// copy[0] = 100

// fmt.Println("Original:", original)
// fmt.Println("Copy:", copy)

// fmt.Println(modifyArray(original))
// fmt.Println(original)

// Practice Ex
// daily_steps := [7]int{1000, 10000, 25543, 2312, 1000, 25000, 3465}

// steps_sum := 0
// for _, step := range daily_steps {
// 	steps_sum += step
// }
// fmt.Println("Sum of steps during the week:", steps_sum)

// max_steps_amount := daily_steps[0]
// max_step_index := 0
// for index, step := range daily_steps[1:] {
// 	if step > max_steps_amount {
// 		max_steps_amount = step
// 		max_step_index = index + 1
// 	}
// }

// fmt.Printf("Max days step count %d\nThe day is %d\n", max_steps_amount, max_step_index)
// average_steps := steps_sum / len(daily_steps)
// fmt.Println("Average steps per day is:", average_steps)

// Homework
// 	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
// 	expenses := [7]float64{12.0, 24.0, 32.93, 120.0, 234.21, 10.0, 30.43}
// 	for index, day := range days {
// 		fmt.Printf("%s: %.2f$.\n", day, expenses[index])
// 	}

// 	total_expenses := 0.0
// 	for _, expense := range expenses {
// 		total_expenses += expense
// 	}
// 	fmt.Printf("Total expensies: %.2f$.\n", total_expenses)

// 	average_expenses := total_expenses / float64(len(expenses))
// 	fmt.Printf("Average expenses: %.2f$.\n", average_expenses)

// 	highest_expense := expenses[0]
// 	highest_expense_day := days[0]
// 	lowest_expense := expenses[0]
// 	lowest_expense_day := days[0]

// 	for day, expense := range expenses[1:] {
// 		if expense > highest_expense {
// 			highest_expense = expense
// 			highest_expense_day = days[day+1]
// 		}
// 		if expense < lowest_expense {
// 			lowest_expense = expense
// 			lowest_expense_day = days[day+1]
// 		}
// 	}

// 	fmt.Printf("The lowest expense day was %s. Expenses was %.2f$.\n", lowest_expense_day, lowest_expense)
// 	fmt.Printf("The highest expense day was %s. Expenses was %.2f$.\n", highest_expense_day, highest_expense)

// 	spend_more_that_ave := 0
// 	for _, expense := range expenses {
// 		if expense > average_expenses {
// 			spend_more_that_ave++
// 		}
// 	}

// 	fmt.Printf("You spend more that average %d times.\n", spend_more_that_ave)

// }

// func modifyArray(arr [3]int) [3]int {
// 	arr[0] = 999 // This modify the COPY, not the original
// 	return arr
// }
