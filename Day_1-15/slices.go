package main

import (
	"fmt"
	"slices"
)

func main() {
	// === Creating Slices ===
	// // Method 1: Slice literal (most common)
	// numbers := []int{1, 2, 3, 4, 5}
	// fmt.Println("Slice literal:", numbers)

	// // Method 2: From an array
	// array := [5]int{10, 20, 30, 40, 50}
	// slice := array[1:4] // element with index 1, 2, 3. NOT 4!
	// fmt.Println("From array:", slice)

	// // Method 3: Using make (create empty slice with specific length)
	// // Syntax: make([]Type, length, capacity)
	// emptySlice := make([]int, 5)
	// fmt.Println("Using make:", emptySlice)

	// // Method 4: Empty slice (length 0)
	// var empty []int
	// fmt.Println("Empty slice:", empty)
	// fmt.Println("Is nil?", empty == nil)

	// === Accessing Slice Elements ===
	// fruits := []string{"apple", "banana", "cherry", "date"}

	// // Access elements (same as array)
	// fmt.Println("First fruit:", fruits[0])
	// fmt.Println("Last fruit:", fruits[len(fruits)-1])

	// // len()
	// fmt.Println("Number of fruits:", len(fruits))

	// // Modify elements
	// fruits[len(fruits)-1] = "blueberry"
	// fmt.Println("Modified:", fruits)

	// // Loop through slice (same as array)
	// fmt.Println("\nAll fruits:")
	// for index, fruit := range fruits {
	// 	fmt.Printf("%d: %s\n", index, fruit)
	// }

	// === Appending to Slice ===
	// Start with an empty slice
	// numbers := []int{}
	// fmt.Println("Initial:", numbers)
	// fmt.Println("Length:", len(numbers))

	// numbers = append(numbers, 1)
	// fmt.Printf("Pointer: %p After append(1): %v\n", numbers, numbers)

	// numbers = append(numbers, 2)
	// fmt.Printf("Pointer: %p After append(2): %v\n", numbers, numbers)

	// numbers = append(numbers, 3)
	// fmt.Println("After append(3):", numbers)

	// // appand multiple elements at once
	// numbers = append(numbers, 4, 5, 6)
	// fmt.Println("After append(4, 5, 6):", numbers)

	// numbers = append(numbers, 7)
	// fmt.Printf("Pointer: %p Final: %v\n", numbers, numbers)

	// === Slicing Operations ===
	// numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println("Original:", numbers)

	// // Syntax: slice[start:end] - include start, EXCLUDE end [start, end)

	// // Get element from index 2 to 5
	// subset := numbers[2:5]
	// fmt.Println("numbers[2:5]:", subset)

	// // From start to index 4
	// start := numbers[:4]
	// fmt.Println("numbers[:4]:", start)

	// // From index 6 to end
	// end := numbers[6:]
	// fmt.Println("numbers[6:]", end)

	// // All elements (create a new view of the slice)
	// all := numbers[:]
	// fmt.Println("numbers[:]:", all)

	// // Slicing create a view NOT A COPY!
	// // Changes to the slice affect the original
	// subset[0] = 999
	// fmt.Println("Modifier subset:", subset)
	// fmt.Println("Original now:", numbers)

	// === Lenght vs Capacity ===
	// numbers := []int{1, 2, 3}

	// // len() - umber of elements currently in the slice
	// fmt.Println("Lenght:", len(numbers))

	// // cap() - total space available (capacity)
	// // Capacity = how many elements the slice CAN hold without reallocating
	// fmt.Println("Capacity:", cap(numbers))

	// // Add more elements
	// numbers = append(numbers, 4)
	// fmt.Println("\nAfter append(4):")
	// fmt.Println("Slice:", numbers)
	// fmt.Println("Lenght:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	// // Add more
	// numbers = append(numbers, 5, 6, 7)
	// fmt.Println("\nAfter append(5,6,7):")
	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	// === Using make() with length and capacity ===
	// numbers := make([]int, 3, 5)
	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	// // First 3 slots are filled with zeros
	// // 2 more slots are available for growth

	// numbers = append(numbers, 10)
	// fmt.Println("\nAfter append(10):")
	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	// numbers = append(numbers, 20)
	// fmt.Println("\nAfter append(20):")
	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	// numbers = append(numbers, 30)
	// fmt.Println("\nAfter append(30):")
	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	// === Copying slices ===
	// original := []int{1, 2, 3, 4, 5}

	// // Create a VIEW
	// view := original[:]
	// view[0] = 999
	// fmt.Println("Original:", original)
	// fmt.Println("View:", view)

	// // Reset
	// original = []int{1, 2, 3, 4, 5}

	// // Use copy() - copy(destination, source)
	// realCopy := make([]int, len(original)) // create empty slice same size
	// copy(realCopy, original)

	// realCopy[0] = 999
	// fmt.Println("\nOriginal:", original)
	// fmt.Println("Copy:", realCopy)

	// small := make([]int, 3)
	// copied := copy(small, original) // Only copied 3 elements
	// fmt.Println("\nCopied", copied, "elements:", small)

	// === Removing elements from Slices ===
	// numbers := []int{10, 20, 30, 40, 50}
	// fmt.Println("Original:", numbers)

	// // Remove element at index 2 (value 30)
	// // Strategy: concatenate slice before and after the element
	// indexToRemove := 2
	// numbers = append(numbers[:indexToRemove], numbers[indexToRemove+1:]...)
	// fmt.Println("After removing index 2:", numbers)

	// Example of ... (spread operator)
	// fmt.Println("\nDemonstrating ... operator")
	// a := []int{1, 2, 3}
	// b := []int{4, 5, 6}
	// combine := append(a, b...) // Unpacks b
	// fmt.Println("Combined:", combine)

	// === Common slice pattern ===
	//// Pattern 1: Building a slice with append
	//var numbers []int // Empty slice
	//for i := 1; i <= 5; i++ {
	//	numbers = append(numbers, i*10)
	//}
	//fmt.Println("Built slice:", numbers)
	//
	//// Pattern 2: Filter (keep only even numbers)
	//filtered := []int{}
	//for _, num := range numbers {
	//	if num%20 == 0 {
	//		filtered = append(filtered, num)
	//	}
	//}
	//fmt.Println("Filtered:", filtered)
	//
	//// Pattern 3: Check if slice contains a value
	//target := 30
	//found := false
	//for _, num := range numbers {
	//	if num == target {
	//		found = true
	//		break
	//	}
	//}
	//fmt.Printf("Contains %d: %v\n", target, found)
	//
	//// Pattern 4: Reverse a slice
	//reversed := []int{}
	//for i := len(numbers) - 1; i >= 0; i-- {
	//	reversed = append(reversed, numbers[i])
	//}
	//fmt.Println("Reversed:", reversed)
	//
	//// Pattern 5: Sun all elements
	//sum := 0
	//for _, num := range numbers {
	//	sum += num
	//}
	//fmt.Println("Sum:", sum)

	// Practice
	//empty := []int{}
	//for i := 1; i <= 10; i++ {
	//	empty = append(empty, i)
	//}
	//fmt.Println("empty:", empty)
	//
	//onlyEven := []int{}
	//for _, num := range empty {
	//	if num%2 == 0 {
	//		onlyEven = append(onlyEven, num)
	//	}
	//}
	//fmt.Println("onlyEven:", onlyEven)
	//
	//sum := 0
	//for _, num := range empty {
	//	sum += num
	//}
	//fmt.Println("sum:", sum)

	// Homework
	studentNames := []string{}
	studentGrades := []float64{}
	userChoice := 0

	for userChoice != 5 {
		var isDataOk bool
		var name string
		var grade float64

		userChoice = 0
		fmt.Println("=== Student Grade Manager ===")
		fmt.Println("1. Add student\n" +
			"2. Display all students\n" +
			"3. Show statistics\n" +
			"4. Remove student\n" +
			"5. Exit\n")
		fmt.Print("Choice: ")
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			isDataOk, name, grade = getStudentData()
			if isDataOk {
				studentNames = addStudentName(studentNames, name)
				studentGrades = addStudentGrade(studentGrades, grade)
			}
		case 2:
			showAllStudents(studentNames, studentGrades)
		case 3:
			getStatistics(studentNames, studentGrades)
		case 4:
			isDataOk, name, grade = getStudentData()
			if isDataOk {
				index := isStudentExist(studentNames, studentGrades, name, grade)
				if index >= 0 {
					studentNames, studentGrades = deleteStudent(studentNames, studentGrades, index)
				} else {
					fmt.Println("Student does not exist")
				}
			}
		}
	}
}

func getStudentData() (bool, string, float64) {
	name := ""
	fmt.Println("Print student name: ")
	fmt.Scan(&name)

	fmt.Println("Print grade: ")
	grade := 0.0
	fmt.Scan(&grade)

	if name == "" && grade < 0 {
		fmt.Println("No student name")
		fmt.Println("Invalid grade")
		return false, "", 0.0
	} else if name == "" {
		fmt.Println("No student name")
		return false, "", grade
	} else if grade < 0 {
		fmt.Println("Invalid grade")
		return false, name, grade
	} else {
		return true, name, grade
	}
}

func addStudentName(studentNames []string, name string) []string {
	studentNames = append(studentNames, name)
	return studentNames
}

func addStudentGrade(studentGrades []float64, grade float64) []float64 {
	studentGrades = append(studentGrades, grade)
	return studentGrades
}

func showAllStudents(studentNames []string, studentGrades []float64) {
	for index, name := range studentNames {
		fmt.Printf("Student name: %s, grade: %f\n", name, studentGrades[index])
	}
}

func getAverage(studentGrades []float64) float64 {
	average := 0.0
	for _, grade := range studentGrades {
		average += grade
	}
	average /= float64(len(studentGrades))
	return average
}

func getHighestGrade(studentGrades []float64) (float64, int) {
	highestGrade := 0.0
	indexOfGrade := 0
	for index, grade := range studentGrades {
		if grade > highestGrade {
			highestGrade = grade
			indexOfGrade = index
		}
	}
	return highestGrade, indexOfGrade
}

func getLowestGrade(studentGrades []float64) (float64, int) {
	lowestGrade := 101.0
	indexOfGrade := 0
	for index, grade := range studentGrades {
		if grade < lowestGrade {
			lowestGrade = grade
			indexOfGrade = index
		}
	}
	return lowestGrade, indexOfGrade
}

func getStatistics(studentNames []string, studentGrades []float64) {
	classAverage := getAverage(studentGrades)
	fmt.Println("ClassAverage =", classAverage)

	fmt.Println("Students above average grade:")
	for index, grade := range studentGrades {
		if grade > classAverage {
			fmt.Printf("Student name: %s; Student grade: %f\n", studentNames[index], grade)
		}
	}

	fmt.Println("Students below average grade:")
	for index, grade := range studentGrades {
		if grade < classAverage {
			fmt.Printf("Student name: %s; Student grade: %f\n", studentNames[index], grade)
		}
	}

	highestGrade, studentHighestGradeIndex := getHighestGrade(studentGrades)
	fmt.Printf("The Highest grade %f has student %s\n", highestGrade, studentNames[studentHighestGradeIndex])

	lowestGrade, studentLowestGradeIndex := getLowestGrade(studentGrades)
	fmt.Printf("The Lowest grade %f has student %s\n", lowestGrade, studentNames[studentLowestGradeIndex])

}

func isStudentExist(studentNames []string, studentGrades []float64, name string, grade float64) int {
	studentNameIndex := slices.Index(studentNames, name)
	if studentNameIndex == -1 {
		return -1
	}
	studentGradeIndex := slices.Index(studentGrades, grade)
	if studentGradeIndex == -1 {
		return -1
	}
	if studentNameIndex == studentGradeIndex {
		return studentNameIndex
	} else {
		return -1
	}
}

func deleteStudent(studentNames []string, studentGrades []float64, index int) ([]string, []float64) {
	studentNames = append(studentNames[:index], studentNames[index+1:]...)
	studentGrades = append(studentGrades[:index], studentGrades[index+1:]...)
	fmt.Println("Student Grades:", studentGrades)
	fmt.Println("Student Names:", studentNames)
	return studentNames, studentGrades
}
