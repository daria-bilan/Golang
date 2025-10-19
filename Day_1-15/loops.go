package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"strings"
// )

// func main() {
// // Basic for loop
// fmt.Println("Counting from 1 to 5:")
// for i := 1; i <= 5; i++ {
// 	fmt.Println(i)
// }

// // Counting backwards
// fmt.Println("\nCountdown:")
// for i := 5; i >= 1; i-- {
// 	fmt.Println(i)
// }

// // Step by 2
// fmt.Println("\nEven numbers 0 to 10:")
// for i := 0; i <= 10; i += 2 {
// 	fmt.Println(i)
// }

// For loop acting as while loop
// counter := 1
// fmt.Println("While-style loop:")
// for counter <= 5 {
// 	fmt.Println(counter)
// 	counter++
// }

// // Practical example: Reading until condition
// sum := 0
// num := 1

// fmt.Println("\nSum until >= 100:")
// for sum < 100 {
// 	sum += num
// 	fmt.Printf("Added %d, ssum is now  %d\n", num, sum)
// 	num++
// }

// Infinite loop
// counter := 0

// for {
// 	counter++
// 	fmt.Println("Iteration:", counter)

// 	if counter >= 5 {
// 		break
// 	}
// }

// numbers := [5]int{10, 20, 30, 40, 50}

// fmt.Println("With index and value:")
// for index, value := range numbers {
// 	fmt.Printf("Index %d: Value%d\n", index, value)
// }

// // Only value (ignore index with _)
// fmt.Println("\nOnly values:")
// for _, value := range numbers {
// 	fmt.Println(value)
// }

// // Only index (just omit the value)
// fmt.Println("\nOnly indices:")
// for index := range numbers {
// 	fmt.Println(index)
// }

// text := "Hello"

// // Range over string gives index and rune
// fmt.Println("Iterating over string:")
// for index, char := range text {
// 	fmt.Printf("Index %d: %c (Unicode: %d)\n", index, char, char)
// }

// ukrainianText := "Привіт"

// fmt.Println("\nUkrainian text:")
// for index, char := range ukrainianText {
// 	fmt.Printf("Index %d: %c\n", index, char)
// }

// Range Loop with Maps
// countries := map[string]string{
// 	"UA": "Ukraine",
// 	"US": "United States",
// 	"UK": "United Kingdom",
// 	"DE": "Germany",
// }

// // Iterate over map
// fmt.Println("Countries:")
// for code, name := range countries {
// 	fmt.Printf("%s: %s\n", code, name)
// }

// // Only keys
// fmt.Println("\nOnly codes:")
// for code := range countries {
// 	fmt.Println(code)
// }

// // Only value
// fmt.Println("\nOnly name:")
// for _, name := range countries {
// 	fmt.Println(name)
// }

// Find first number divisible by 7
// fmt.Println("Finding first number divisible by 7:")
// for i := 1; i <= 100; i++ {
// 	if i%7 == 0 {
// 		fmt.Printf("Found: %d\n", i)
// 		break
// 	}
// }

// Search in array
// numbers := [10]int{3, 7, 12, 5, 18, 21, 9, 15, 6, 11}
// target := 13
// found := false

// fmt.Printf("\nSearching for %d:\n", target)
// for index, num := range numbers {
// 	if num == target {
// 		fmt.Printf("Found %d at index %d\n", target, index)
// 		found = true
// 		break
// 	}

// }

// if !found {
// 	fmt.Printf("%d not found\n", target)
// }

// continue statement
// fmt.Println("Odd numbers from 1 to 10:")
// for i := 1; i <= 10; i++ {
// 	if i%2 == 0 {
// 		continue
// 	}
// 	fmt.Println(i)
// }

// numbers := [8]int{5, -3, 8, -1, 12, -7, 4, 9}
// sum := 0

// fmt.Println("\nSum of positive numbers:")
// for _, num := range numbers {
// 	if num < 0 {
// 		fmt.Printf("Skipping %d\n", num)
// 		continue
// 	}
// 	sum += num
// 	fmt.Printf("Adding %d, sum = %d\n", num, sum)
// }

// fmt.Printf("\nFinal sum: %d\n", sum)

// Nested loops
// fmt.Println("Multiplication Table (5x5):")
// for i := 1; i <= 5; i++ {
// 	for j := 1; j <= 5; j++ {
// 		fmt.Printf("%4d", i*j)
// 	}
// 	fmt.Println()
// }

// fmt.Println("\nTriangle pattern:")
// for i := 1; i <= 5; i++ {
// 	for j := 1; j <= i; j++ {
// 		fmt.Print("* ")
// 	}
// 	fmt.Println()
// }

// 	fmt.Println("Without label:")
// 	for i := 1; i <= 3; i++ {
// 		for j := 1; j <= 3; j++ {
// 			if j == 2 {
// 				break
// 			}
// 			fmt.Printf("i=%d, j=%d\n", i, j)
// 		}
// 	}

// 	fmt.Println("\nWith label:")
// outer:
// 	for i := 1; i <= 3; i++ {
// 		for j := 1; j <= 3; j++ {
// 			if j == 2 {
// 				break outer
// 			}
// 			fmt.Printf("i=%d, j=%d\n", i, j)
// 		}
// 	}

// 	fmt.Println("Done!")

// Practice
// number := 7
// fmt.Println("Times table for number:", number)
// for i := 1; i <= 10; i++ {
// 	fmt.Printf("%d x %d = %d\n", number, i, number*i)
// }

// fmt.Println("\nAll prime numbers between 1 and 50:")
// isPrime := true
// for num := 2; num <= 50; num++ {
// 	for i := 2; i < num; i++ {
// 		if num%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}
// 	if isPrime {
// 		fmt.Printf("%d ", num)
// 	}
// }
// fmt.Println()

// fmt.Println("Checkerboard pattern (8x8)")
// for i := 0; i < 8; i++ {
// 	for j := 0; j < 8; j++ {
// 		if (i+j)%2 == 0 {
// 			fmt.Print("■ ")
// 		} else {
// 			fmt.Print("□ ")
// 		}
// 	}
// 	fmt.Println()
// }

// Homework

// newGame:
// 	for {
// 		var answer int
// 		isGuessed := false
// 		target := rand.Intn(100) + 1
// 		attempt := 0
// 		lives := 7
// 	currentSession:
// 		for lives > 0 {
// 			fmt.Printf("You have %d attempts left.\n", lives)
// 			fmt.Println("Guess a number between 1 and 100: ")
// 			fmt.Scan(&answer)
// 			if answer == target {
// 				fmt.Println("Correct!")
// 				attempt++
// 				isGuessed = true
// 			} else if answer == -1 {
// 				break newGame
// 			} else if answer > target {
// 				fmt.Println("Too high!")
// 				attempt++
// 				lives--
// 			} else {
// 				fmt.Println("Too low!")
// 				attempt++
// 				lives--
// 			}

// 			if lives == 0 {
// 				fmt.Println("Game Over! You lost all your attempts and didn't guess a number!")
// 				fmt.Printf("The guessed number was %d\n", target)
// 			}

// 			if isGuessed {
// 				fmt.Printf("Game Over!\nYou guessed in %d attempts.\n", attempt)
// 				break currentSession
// 			}

// 		}
// 		var isContinue string
// 		fmt.Println("Do you want restart a game?")
// 		fmt.Scan(&isContinue)
// 		if strings.ToLower(isContinue) == "yes" {
// 			continue
// 		} else if strings.ToLower(isContinue) == "no" {
// 			fmt.Printf("Game Over!\nYou guessed in %d attempts.\n", attempt)
// 			break newGame
// 		}
// 	}
// }
