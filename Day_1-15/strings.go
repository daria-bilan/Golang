package main

func main() {
	//// 1. Basic strings operation
	//// Define a sample string
	//text := "Hello, Go World!"
	//
	//// Length of string (in bytes)
	//fmt.Printf("Length of '%s': %d bytes\n", text, len(text))
	//
	//// Convert to uppercase and lowercase
	//upper := strings.ToUpper(text)
	//lower := strings.ToLower(text)
	//fmt.Printf("Uppercase: %s\nLowercase: %s\n", upper, lower)
	//
	////Trim whitespace or specific characters
	//trimmedSpace := strings.TrimSpace("   Hello    \n\t!")
	//fmt.Printf("Trimmed whitespaces: %s\n", trimmedSpace)
	//
	//trimmedChars := strings.Trim("**Hello**!", "*")
	//fmt.Printf("Trimmed special characters: %s\n", trimmedChars)
	//
	////Split strings into slice
	//words := strings.Split(text, " ")
	//fmt.Printf("Words: %v\n", words)
	//
	////Check for substrings
	//hasGo := strings.Contains(text, "Go")
	//fmt.Printf("Contains 'Go': %t\n", hasGo)

	//// 2. Working with Runes
	//// String with Unicode characters
	//text := "Hello😊World!"
	//
	//// Length in bytes vs runes
	//fmt.Printf("Bytes: %d\n", len(text))
	//runes := []rune(text)
	//fmt.Printf("Runes: %d\n", len(runes))
	//
	//// Iterate over runes to process characters
	//for i, r := range runes {
	//	fmt.Printf("Index %d: %c (Unicode U+%04X)\n", i, r, r)
	//}

	//// 3. Practical Text Processing
	//input := "  Hello   GO  World!  "
	//result := cleanInput(input)
	//fmt.Printf("Original: '%s'\nCleaned: '%s'\n", input, result)

	// Practice exercises

}

//func cleanInput(input string) string {
//	//Remove leading/trailing space
//	cleaned := strings.TrimSpace(input)
//
//	// Convert to lowercase for case-intensitive comparison
//	cleaned = strings.ToLower(cleaned)
//
//	// Replace multiple spaces with single space
//	// - strings.Replace: Replace up to n occurrences
//	replaced := strings.Replace(cleaned, "  ", " ", 1) // Replace first 2 space "  " to 1 space " "
//	fmt.Printf("After Replace(1): '%s'\n", replaced)
//	//- strings.ReplaceAll: Replace all occurrences
//	cleaned = strings.ReplaceAll(cleaned, "  ", " ")
//	return cleaned
//}
