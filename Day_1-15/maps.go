package main

func main() {
	// Creating maps
	//// Method 1: Map literal (most common)
	//ages := map[string]int{
	//	"Alice":   25,
	//	"Bob":     30,
	//	"Charlie": 35,
	//}
	//fmt.Println("Map literal:", ages)
	//
	//// Method 2: Using make
	//// 🔮 make() creates an empty map ready to use
	//// Syntax: make(map[KeyType]ValueType)
	//scores := make(map[string]int)
	//fmt.Println("Using make:", scores)
	//
	//// Method 3: Empty map declaration
	//var colors map[string]string
	//fmt.Println("Empty declaration:", colors)
	//fmt.Println("Is nil?", colors == nil)
	//
	//// ⚠️ IMPORTANT: nil map vs empty map
	//// nil map: cannot add items (causes panic!)
	//// empty map (from make): can add items
	//
	//// This would PANIC:
	//// colors["sky"] = "blue"  // ❌ panic: assignment to entry in nil map
	//// Must initialize first
	//colors = make(map[string]string)
	//colors["sky"] = "blue"
	//fmt.Println("After initialization:", colors)

	//// Map syntax and type
	//// Syntax: map[KeyType]ValueType
	//
	//// Common map type
	//map[string]int         // String key -> Integer value
	//map[string]string      // String key -> String value
	//map[int]string         // Integer key -> String value
	//map[string]float64     // String key -> Float value
	//map[string]bool        // String key -> Boolean value

	////Examples:
	//ages := map[string]int{
	//	"Alice": 25,
	//	"Bob":   30,
	//}
	//
	//capitals := map[string]string{
	//	"Ukraine": "Kyiv",
	//	"France":  "Paris",
	//	"Japan":   "Tokyo",
	//}
	//
	//studentPassed := map[string]bool{
	//	"Alice":   true,
	//	"Bob":     false,
	//	"Charlie": true,
	//}
	//
	//// ✅ Valid key types
	//map[string]int       // String keys
	//map[int]string       // Integer keys
	//map[float64]bool     // Float keys (but be careful with precision!)
	//
	//// ❌ Invalid key types
	//// map[[]int]string  // ❌ Slice keys - not allowed!
	//// map[map[string]int]string  // ❌ Map keys - not allowed!

	//// Adding and Accessing Value
	//grade := make(map[string]int)
	//
	//// Add values
	//grade["Alice"] = 95
	//grade["Bob"] = 100
	//grade["Carol"] = 87
	//
	//fmt.Println("Grades:", grade)
	//
	//// Access values
	//aliceGrade := grade["Alice"]
	//fmt.Println("Alice grade:", aliceGrade)
	//
	//// Update values
	//grade["Bob"] = 89
	//fmt.Println("Bob's new grade:", grade["Bob"])
	//
	//// Access non-existent key
	//// Returns zero value for value type
	//davidGrade := grade["David"]
	//fmt.Println("David grade:", davidGrade)

	//⚠️ Important: Accessing a non-existent key returns the zero value for the value type:
	//
	//map[string]int → missing key returns 0
	//map[string]string → missing key returns ""
	//map[string]bool → missing key returns false

	////  Checking If Key Exists (The Comma-OK Idiom)
	//grades := map[string]int{
	//	"Alice":   95,
	//	"Bob":     87,
	//	"Charlie": 0,
	//}
	//
	//// Check Alice exist
	//aliceGrade, ok := grades["Alice"]
	//if ok {
	//	fmt.Printf("Alice exist with grade: %d\n", aliceGrade)
	//} else {
	//	fmt.Println("Alice does not exist")
	//}
	//
	//// Check Charlie (exist with grade 0)
	//charlieGrade, ok := grades["Charlie"]
	//if ok {
	//	fmt.Printf("Charlie exist with grade: %d\n", charlieGrade)
	//} else {
	//	fmt.Println("Charlie does not exist")
	//}
	//
	//// Check David (doesn't exist)
	//davidGrade, ok := grades["David"]
	//if ok {
	//	fmt.Printf("David exist with grade: %d\n", davidGrade)
	//} else {
	//	fmt.Println("David does not exist")
	//}
	//
	//// Common pattern: check and use in one line
	//if grade, exist := grades["Bob"]; exist {
	//	fmt.Printf("Bob's grade: %d\n", grade)
	//}
	//// Just checking existence (ignore value with _)
	//if _, exist := grades["Alice"]; exist {
	//	fmt.Println("Alice is in the map")
	//}

	//// Deleting Entries
	//ages := map[string]int{
	//	"Alice":   25,
	//	"Bob":     30,
	//	"Charlie": 35,
	//	"Davis":   28,
	//}
	//
	//fmt.Println("Before deleting:", ages)
	//
	//// Delete a key using built-in delete() function
	//// 🔮 delete(map, key) removes the key-value pair
	//// Syntax: delete(mapName, key)
	//
	//delete(ages, "Bob")
	//fmt.Println("After deleting Bob:", ages)
	//
	//// Deleting non-existence key is safe (no error)
	//delete(ages, "Eve")
	//fmt.Println("After deleting Eve (doesn't exist):", ages)
	//
	//// Check if Bob still exists
	//if _, exist := ages["Bob"]; exist {
	//	fmt.Println("Bob exist")
	//} else {
	//	fmt.Println("Bob was deleted")
	//}

	//// Iterating Over Maps
	//// ⚠️ IMPORTANT: Map iteration order is RANDOM!
	//capitals := map[string]string{
	//	"Ukraine": "Kyiv",
	//	"France":  "Paris",
	//	"Japan":   "Tokyo",
	//	"USA":     "Washington",
	//}
	//// Method 1: Iterate over key-value pairs
	//fmt.Println("Countries and capitals:")
	//for country, capital := range capitals {
	//	fmt.Printf("%s -> %s\n", country, capital)
	//}
	//
	//// Method 2: Iterate over keys only
	//fmt.Println("\nCountries only:")
	//for country := range capitals {
	//	fmt.Println(country)
	//}
	//
	//// Method 3: Iterate over values only
	//fmt.Println("\nOnly capitals:")
	//for _, capital := range capitals {
	//	fmt.Println(capital)
	//}

	//// If you need sorted output
	//countries := []string{}
	//for country := range capitals {
	//	countries = append(countries, country)
	//}
	//
	//// Sort the slice
	//sort.Strings(countries)
	//
	//// Now iterate in sorted order
	//for _, country := range countries {
	//	fmt.Printf("%s -> %s\n", country, capitals[country])
	//}

	//// Map Length
	//scores := map[string]int{
	//	"Alice":   95,
	//	"Bob":     87,
	//	"Charlie": 92,
	//}
	//
	//// Get number of key-value pairs using len()
	//count := len(scores)
	//fmt.Printf("Numer of students: %d\n", count)
	//
	//// Add more entries
	//scores["David"] = 88
	//scores["Eve"] = 91
	//
	//fmt.Printf("After adding 2 students: %d\n", len(scores))
	//
	//// Delete entries
	//delete(scores, "Bob")
	//
	//fmt.Printf("After removing Bob: %d\n", len(scores))
	//
	//// Empty map
	//empty := make(map[string]int)
	//fmt.Printf("Empty map length: %d\n", len(empty))

	//// Maps and Sets
	//// Use map[Type]bool to create a set
	//// The bool value doesn't matter, only key existence matters
	//uniqueWord := make(map[string]bool)
	//
	//// Add items to set (value = true)
	//uniqueWord["apple"] = true
	//uniqueWord["banana"] = true
	//uniqueWord["cherry"] = true
	//uniqueWord["apple"] = true
	//
	//fmt.Println("Unique words:", uniqueWord)
	//fmt.Println("Number of unique words:", len(uniqueWord))
	//
	//// Check if item is in set
	//if uniqueWord["apple"] {
	//	fmt.Println("apple is in the set")
	//}
	//if uniqueWord["grape"] {
	//	fmt.Println("grape is in the set")
	//} else {
	//	fmt.Println("grape is NOT in the set")
	//}
	//
	//// Better check using comma-ok
	//if _, exist := uniqueWord["banana"]; exist {
	//	fmt.Println("banana is in the set")
	//}
	//
	//// Remove from set
	//delete(uniqueWord, "apple")
	//
	//// Get all items in set
	//fmt.Println("\nAll items is set:")
	//for word := range uniqueWord {
	//	fmt.Println(word)
	//}

	//// Map with Complex Values
	//// Map with slice values
	//// Each country can have multiple cities
	//cities := map[string][]string{
	//	"Ukraine": {"Kyiv", "Lviv", "Odesa"},
	//	"France":  {"Paris", "Lyon", "Marseille"},
	//	"Japan":   {"Tokyo", "Osaka", "Kyoto"},
	//}
	//
	//fmt.Println("Cities in Ukraine:", cities["Ukraine"])
	//
	//// Add a city to Ukraine
	//cities["Ukraine"] = append(cities["Ukraine"], "Kharkiv")
	//fmt.Println("Updated Ukraine cities:", cities["Ukraine"])
	//
	//// Add a new country
	//cities["USA"] = []string{"New York", "Los Angeles"}
	//
	//// Iterate and print all
	//fmt.Println("\nAll cities by country:")
	//for country, cityList := range cities {
	//	fmt.Printf("%s: ", country)
	//	for i, city := range cityList {
	//		if i > 0 {
	//			fmt.Print(", ")
	//		}
	//		fmt.Print(city)
	//	}
	//	fmt.Println()
	//}

	//// Nested map
	//// Student -> Subject -> Grade
	//studentGrades := map[string]map[string]int{
	//	"Alice": {
	//		"Math":    95,
	//		"English": 88,
	//		"Science": 92,
	//	},
	//	"Bob": {
	//		"Math":    87,
	//		"English": 90,
	//		"Science": 85,
	//	},
	//}
	//
	//fmt.Println("\nAlice's Math grade:", studentGrades["Alice"]["Math"])
	//
	//// Add new subject for Alice
	//studentGrades["Alice"]["History"] = 91
	//
	//// Add new student
	//studentGrades["Charlie"] = make(map[string]int)
	//studentGrades["Charlie"]["Math"] = 89
	//studentGrades["Charlie"]["English"] = 93
	//
	//// Print all grades
	//fmt.Println("\nAll student grades:")
	//for student, subjects := range studentGrades {
	//	fmt.Printf("%s:\n", student)
	//	for subject, grade := range subjects {
	//		fmt.Printf("	%s: %d\n", subject, grade)
	//	}
	//}

	// Common Map Patterns
	//// Pattern 1: Counting occurrences (frequency map)
	//words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	//
	//wordCount := make(map[string]int)
	//for _, word := range words {
	//	wordCount[word]++
	//}
	//
	//fmt.Println("Word frequencies:", wordCount)

	//// Pattern 2: Grouping by key
	//students := []string{"Alice", "Bob", "Andrew", "Charlie", "Anna"}
	//
	//byFirstLetter := make(map[string][]string)
	//for _, student := range students {
	//	firstLetter := string(student[0])
	//	byFirstLetter[firstLetter] = append(byFirstLetter[firstLetter], student)
	//}
	//
	//fmt.Println("\nStudents grouped by first letter:")
	//for letter, names := range byFirstLetter {
	//	fmt.Printf("%s: %v\n", letter, names)
	//}

	//// Pattern 3: Lookup table (fast checking)
	//validUsers := map[string]bool{
	//	"alice":   true,
	//	"bob":     true,
	//	"charlie": true,
	//}
	//
	//username := "bob"
	//if validUsers[username] {
	//	fmt.Printf("\n%s is a valid user\n", username)
	//}

	//// Pattern 4: Default values
	//settings := map[string]string{
	//	"theme": "dark",
	//	"lang":  "en",
	//}
	//
	//// Get value with default if not exist
	//theme := settings["theme"]
	//if theme == "" {
	//	theme = "light" // Default
	//}
	//fmt.Printf("\nTheme: %s\n", theme)
	//// Or using comma-ok
	//fontSize, exist := settings["fontSize"]
	//if !exist {
	//	fontSize = "12px"
	//}
	//fmt.Printf("Font size: %s\n", fontSize)

	//// Pattern 5: Inverting a map (swap keys and values)
	//original := map[string]string{
	//	"Ukraine": "Kyiv",
	//	"France":  "Paris",
	//	"Japan":   "Tokyo",
	//}
	//
	//inverted := make(map[string]string)
	//for country, capital := range original {
	//	inverted[capital] = country
	//}
	//
	//fmt.Println("\nInverted map (capital -> country):", inverted)

	//// Pattern 6: Merging maps
	//map1 := map[string]int{"a": 1, "b": 2}
	//map2 := map[string]int{"b": 3, "c": 4}
	//
	//merged := make(map[string]int)
	//
	//// Copy map1
	//for k, v := range map1 {
	//	merged[k] = v
	//}
	//
	//// Add map2 (overwrites duplicates)
	//for k, v := range map2 {
	//	merged[k] = v
	//}
	//
	//fmt.Println("\nMerged map:", merged) // b was overwritten

	// === Practice Exercise ===
	//products := map[string]float64{
	//	"Monitor": 20.32,
	//	"Laptop":  100.00,
	//	"Mouse":   8.15,
	//}
	//
	//fmt.Println("All products:", products)
	//
	//// Add new product
	//products["Lamp"] = 6.12
	//fmt.Println("After adding new product:", products)
	//
	//// Update price
	//products["Laptop"] = 85.99
	//fmt.Println("After updating Laptop price:", products)
	//
	//// Delete product
	//delete(products, "Mouse")
	//fmt.Println("After removing Mouse:", products)
	//
	//totalValue := 0.0
	//for _, v := range products {
	//	totalValue += v
	//}
	//fmt.Printf("Total value: %2.f\n", totalValue)

}
