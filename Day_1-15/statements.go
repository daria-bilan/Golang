package main

// import "fmt"
// func main() {
// score := 75
// if score >= 90 {
// 	fmt.Println("Excellent!")
// } else if score >= 80 {
// 	fmt.Println("Veru good.")
// } else if score >= 70 {
// 	fmt.Println("Good.")
// } else {
// 	fmt.Println("Needs improvement.")
// }

// rand.Seed(time.Now().UnixNano())

// if num := rand.Intn(100); num < 50 {
// 	fmt.Printf("%d is less than 50\n", num)
// } else {
// 	fmt.Printf("%d is greater that or equal to 50\n", num)
// }

// var number float32
// fmt.Scan(&number)
// if number == 0 {
// 	fmt.Printf("%f is a zero", number)
// } else if number > 0 {
// 	fmt.Printf("%f is a positive number.", number)
// } else {
// 	fmt.Printf("%f is a negative number.", number)
// }

// basic for loop
// for i := 0; i < 5; i++ {
// 	fmt.Println(i)
// }

// for loop with single condition
// i := 0
// for i < 5 {
// 	fmt.Println(i)
// 	i++
// }

// infinite loop
// i := 0
// for {
// 	fmt.Println(i)
// 	i++
// 	if i == 5 {
// 		break
// 	}
// }

// range
// numbers := []int{10, 20, 30, 40, 50}
// for index, value := range numbers {
// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
// }

// if only value needed
// for _, value := range collection {
// }

// var sum int
// for i := 1; i < 100; i++ {
// 	if i%2 != 0 {
// 		continue
// 	}
// 	sum += i
// }
// fmt.Printf("Sum: %d\n", sum)

// swithc expression
// grade := "B"
// switch grade {
// case "A":
// 	fmt.Println("Excellent!")
// case "B":
// 	fmt.Println("Good!")
// case "C":
// 	fmt.Println("Fair.")
// default:
// 	fmt.Println("Need improvement!")
// }

// day := "Wednesday"
// switch day {
// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
// 	fmt.Println("It's a weekday.")
// case "Saturday", "Sunday":
// 	fmt.Println("It's a weekend.")
// default:
// 	fmt.Println("Invalid day.")
// }

// switch with go expression
// score := 85
// switch {
// case score >= 90:
// 	fmt.Println("Excellent!")
// case score >= 80:
// 	fmt.Println("Good!")
// case score >= 70:
// 	fmt.Println("Fair.")
// default:
// 	fmt.Println("Need improvement!")
// }

//fallthrough
// num := 2
// switch num {
// case 1:
// 	fmt.Println("One")
// case 2:
// 	fmt.Println("Two")
// 	fallthrough
// case 3:
// 	fmt.Println("Three")
// default:
// 	fmt.Println("Default")
// }

// defer
// defer fmt.Println("Third")
// defer fmt.Println("Second")
// defer fmt.Println("First")

// fmt.Println("Hello")
// fmt.Println("Apple")
// fmt.Println("Sun")

// 	x := 10
// 	defer fmt.Println("Value of x:", x)

// 	x = 20
// 	fmt.Println("Value of x:", x)

// }
