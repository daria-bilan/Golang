package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	//text := "Hello"
	//
	//// Iterate as BYTES
	//fmt.Println("As bytes:")
	//for i := 0; i < len(text); i++ {
	//	fmt.Printf("Index %d: byte %d ('%c')\n", i, text[i], text[i])
	//}
	//
	//// Iterate as RUNES (characters)
	//fmt.Println("\nAs runes:")
	//for index, char := range text {
	//	fmt.Printf("Index %d: rune %d ('%c')\n", index, char, char)
	//}
	//text := "Hello"
	//
	//// Iterate as BYTES
	//fmt.Println("As bytes:")
	//for i := 0; i < len(text); i++ {
	//	fmt.Printf("Index %d: byte %d ('%c')\n", i, text[i], text[i])
	//}
	//
	//// Iterate as RUNES (characters)
	//fmt.Println("\nAs runes:")
	//for index, char := range text {
	//	fmt.Printf("Index %d: rune %d ('%c')\n", index, char, char)
	//}
	//texts := []string{
	//	"Hello",      // English (1 byte per char)
	//	"Привіт",     // Ukrainian (2 bytes per char)
	//	"こんにちは",     // Japanese (3 bytes per char)
	//	"👋🌍",        // Emojis (4 bytes per char)
	//}
	//
	//for _, text := range texts {
	//	// len() counts bytes
	//	byteCount := len(text)
	//
	//	// 🔮 utf8.RuneCountInString() counts actual characters
	//	// This is from unicode/utf8 package
	//	charCount := utf8.RuneCountInString(text)
	//
	//	fmt.Printf("Text: %s\n", text)
	//	fmt.Printf("  Bytes: %d\n", byteCount)
	//	fmt.Printf("  Characters: %d\n", charCount)
	//	fmt.Println()
	//}
	//text := "  Hello, World!  "
	//
	//// 1. ToUpper / ToLower - Change case
	//fmt.Println("ToUpper:", strings.ToUpper(text))
	//fmt.Println("ToLower:", strings.ToLower(text))
	//
	//// 2. TrimSpace - Remove leading/trailing spaces
	//fmt.Println("TrimSpace:", strings.TrimSpace(text))
	//
	//// 3. Trim - Remove specific characters
	//fmt.Println("Trim '! ':", strings.Trim(text, " !"))
	//
	//// 4. Contains - Check if substring exists
	//fmt.Println("Contains 'World':", strings.Contains(text, "World"))
	//fmt.Println("Contains 'Goodbye':", strings.Contains(text, "Goodbye"))
	//
	//// 5. Index - Find position of substring
	//pos := strings.Index(text, "World")
	//fmt.Println("Index of 'World':", pos)
	//
	//notFound := strings.Index(text, "Goodbye")
	//fmt.Println("Index of 'Goodbye':", notFound)  // -1 means not found
	//
	//// 6. Replace - Replace substring
	//fmt.Println("Replace 'World' with 'Go':",
	//	strings.Replace(text, "World", "Go", 1))  // Replace 1 occurrence
	//
	//fmt.Println("Replace all spaces:",
	//	strings.Replace(text, " ", "_", -1))  // -1 = replace all
	//
	//// 7. Split - Break string into slice
	//sentence := "apple,banana,cherry"
	//fruits := strings.Split(sentence, ",")
	//fmt.Println("Split:", fruits)
	//
	//// 8. Join - Combine slice into string
	//joined := strings.Join(fruits, " | ")
	//fmt.Println("Join:", joined)
	//
	//// 9. HasPrefix / HasSuffix
	//filename := "document.txt"
	//fmt.Println("HasPrefix 'doc':", strings.HasPrefix(filename, "doc"))
	//fmt.Println("HasSuffix '.txt':", strings.HasSuffix(filename, ".txt"))
	//
	//// 10. Count - Count occurrences
	//text2 := "hello hello hello"
	//fmt.Println("Count 'hello':", strings.Count(text2, "hello"))

	//// 1. Repeat - Repeat string N times
	//fmt.Println(strings.Repeat("Go", 3))  // "GoGoGo"
	//fmt.Println(strings.Repeat("-", 20))  // "--------------------"
	//
	//// 2. Fields - Split by whitespace (smart split)
	//text := "Hello   World    from     Go"  // Multiple spaces
	//words := strings.Fields(text)
	//fmt.Println("Fields:", words)  // ["Hello", "World", "from", "Go"]
	//
	//// 3. ReplaceAll - Replace all occurrences
	//oldText := "I love cats. Cats are amazing. Cats!"
	//newText := strings.ReplaceAll(oldText, "Cats", "Dogs")
	//fmt.Println(newText)
	//
	//// 4. Compare - Compare strings (returns -1, 0, or 1)
	//fmt.Println(strings.Compare("apple", "banana"))  // -1 (apple < banana)
	//fmt.Println(strings.Compare("apple", "apple"))   // 0 (equal)
	//fmt.Println(strings.Compare("banana", "apple"))  // 1 (banana > apple)
	//
	//// 5. EqualFold - Case-insensitive comparison
	//fmt.Println(strings.EqualFold("Hello", "hello"))  // true
	//fmt.Println(strings.EqualFold("Go", "GO"))        // true
	//
	//// 6. Title - Convert to title case (deprecated but still works)
	//fmt.Println(strings.Title("hello world"))  // "Hello World"
	//
	//// 7. LastIndex - Find last occurrence
	//path := "/home/user/documents/file.txt"
	//lastSlash := strings.LastIndex(path, "/")
	//filename := path[lastSlash+1:]
	//fmt.Println("Filename:", filename)  // "file.txt"

	//// String to rune slice
	//text := "Привіт"
	//runes := []rune(text)  // Convert string to rune slice
	//
	//fmt.Println("Rune slice:", runes)
	//fmt.Println("Length:", len(runes))  // Now this is character count!
	//
	//// Access individual characters correctly
	//fmt.Printf("First character: '%c'\n", runes[0])
	//fmt.Printf("Last character: '%c'\n", runes[len(runes)-1])
	//
	//// Modify a character
	//runes[0] = 'Б'  // Change П to Б
	//
	//// Convert back to string
	//modified := string(runes)
	//fmt.Println("Modified:", modified)  // "Бривіт"
	//
	//// Build string from runes
	//greeting := []rune{'П', 'р', 'и', 'в', 'і', 'т', '!'}
	//fmt.Println(string(greeting))
	//
	//// Single rune
	//var letter rune = 'A'
	//fmt.Printf("Rune: %c, Value: %d\n", letter, letter)
	//
	//// Unicode emojis
	//emoji := '😀'
	//fmt.Printf("Emoji: %c, Value: %d\n", emoji, emoji)

	//// Example 1: Clean and validate email
	//email := "  User@Example.COM  "
	//cleanEmail := strings.ToLower(strings.TrimSpace(email))
	//fmt.Println("Clean email:", cleanEmail)
	//
	//if strings.Contains(cleanEmail, "@") {
	//	parts := strings.Split(cleanEmail, "@")
	//	fmt.Println("Username:", parts[0])
	//	fmt.Println("Domain:", parts[1])
	//}
	//
	//// Example 2: Parse CSV line
	//csvLine := "John,Doe,30,Engineer"
	//fields := strings.Split(csvLine, ",")
	//fmt.Printf("First Name: %s, Last Name: %s, Age: %s, Job: %s\n",
	//	fields[0], fields[1], fields[2], fields[3])
	//
	//// Example 3: Build a formatted message
	//name := "Alice"
	//score := 95
	//message := fmt.Sprintf("Congratulations, %s! You scored %d points!", name, score)
	//fmt.Println(message)
	//
	//// Example 4: Check file extension
	//filename := "document.PDF"
	//if strings.HasSuffix(strings.ToLower(filename), ".pdf") {
	//	fmt.Println("This is a PDF file")
	//}
	//
	//// Example 5: Count words
	//text := "Go is awesome. Go is fast. Go is simple."
	//wordCount := len(strings.Fields(text))
	//fmt.Println("Word count:", wordCount)
	//
	//// Example 6: Remove all spaces
	//phone := "380 123 45 67"
	//cleanPhone := strings.ReplaceAll(phone, " ", "")
	//fmt.Println("Clean phone:", cleanPhone)
	//
	//// Example 7: Create initials
	//fullName := "Alice Bob Charlie"
	//names := strings.Fields(fullName)
	//var initials string
	//for _, name := range names {
	//	if len(name) > 0 {
	//		initials += string(name[0]) + "."
	//	}
	//}
	//fmt.Println("Initials:", initials)

	//// String Builder
	//// Slow way
	//var result string
	//for i := 0; i < 5; i++ {
	//	result += fmt.Sprintf("Line %d\n", i)
	//}
	//fmt.Println("Slow way: ", result)
	//
	//// Fast way (using strings.Builder)
	//var builder strings.Builder
	//for i := 0; i < 5; i++ {
	//	builder.WriteString(fmt.Sprintf("Line %d\n", i))
	//}
	//fmt.Println("Fast way: ", builder.String())
	//
	//// strings.Builder methods
	//var sb strings.Builder
	//sb.WriteString("Hello")
	//sb.WriteString(" ")
	//sb.WriteString("World")
	//sb.WriteByte('!')
	//sb.WriteRune('😀')
	//
	//fmt.Println(sb.String())

	// Practice exercise
	//input := "Hello sunny fluffy World!"
	//wordsCount := len(strings.Fields(input))
	//fmt.Println("Words in input:", wordsCount)
	//longestWord := ""
	//length := 0
	//for _, word := range strings.Fields(input) {
	//	if len(word) > length {
	//		length = len(word)
	//		longestWord = word
	//	}
	//}
	//fmt.Printf("The longest word is '%s'. It's length id %d\n", longestWord, length)
	//aCount, eCount, iCount, oCount, uCount := strings.Count(input, "a"), strings.Count(input, "e"), strings.Count(input, "i"), strings.Count(input, "o"), strings.Count(input, "u")
	//var result strings.Builder
	//result.WriteString(fmt.Sprintf("a count: %d\n", aCount))
	//result.WriteString(fmt.Sprintf("e count: %d\n", eCount))
	//result.WriteString(fmt.Sprintf("i count: %d\n", iCount))
	//result.WriteString(fmt.Sprintf("o count: %d\n", oCount))
	//result.WriteString(fmt.Sprintf("u count: %d\n", uCount))
	//fmt.Println(result.String())
	//
	//var reversedInput strings.Builder
	//for i := len(input) - 1 - 1; i >= 0; i-- {
	//	reversedInput.WriteString(string(input[i]))
	//}
	//fmt.Println(reversedInput.String())

	// === Homework ===
	var input string
	fmt.Println("Enter a string: ")
	//fmt.Scanf("%q", &input)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	// Statistics:
	runesCount := utf8.RuneCountInString(input)
	runesCountWithLen := len([]rune(input))
	totalWords := len(strings.Fields(input))

	sentences := 0
	for _, char := range input {
		if strings.ContainsRune(".!?", char) {
			sentences++
		}
	}

	var longestWord string
	var length int
	for _, word := range strings.Fields(input) {
		wordLength := utf8.RuneCountInString(word)
		if wordLength > length {
			length = wordLength
			longestWord = word
		}
	}

	wordCount := make(map[string]int)
	for _, word := range strings.Fields(strings.ToLower(input)) {
		wordCount[word]++
	}
	mostCommonWord := ""
	maxCount := 0
	for word, count := range wordCount {
		if count > maxCount {
			maxCount = count
			mostCommonWord = word
		}
	}

	fmt.Println("===Statistics===")
	fmt.Println("Input: ", input)
	fmt.Println("Runes count: ", runesCount)
	fmt.Println("Runes count with len([]rune(input): ", runesCountWithLen)
	fmt.Println("Total words count: ", totalWords)
	fmt.Println("Total sentences: ", sentences)
	fmt.Printf("Longest word: %s. It's length: %d\n", longestWord, length)
	fmt.Printf("Most Common: %s. It meets %d times.\n", mostCommonWord, maxCount)

	// Transforming text
	fmt.Println("\n===Transforming text===")
	fmt.Println("To Upper: ", strings.ToUpper(input))
	fmt.Println("To Lower: ", strings.ToLower(input))
	fmt.Println("To Title:", toTitleCase(input))

	var reversedInput strings.Builder
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		reversedInput.WriteRune(runes[i])
	}
	fmt.Println("Reversed input: ", reversedInput.String())
	cleanedWords := strings.Fields(input)
	cleaned := strings.Join(cleanedWords, " ")
	fmt.Println("Remove extra spaces:", cleaned)

	// ===Search functionality:===
	wordForSearch := longestWord
	wordForCount := "vitae"
	wordForReplace := "et"
	isExist := strings.Index(strings.ToLower(input), strings.ToLower(wordForSearch))
	countOccurrences := strings.Count(strings.ToLower(input), strings.ToLower(wordForCount))
	replacedInput := strings.ReplaceAll(strings.ToLower(input), strings.ToLower(wordForReplace), "**")
	fmt.Println("===Search functionality===")
	if isExist == -1 {
		fmt.Println("Not found: ", wordForSearch)
	} else {
		fmt.Printf("Found: '%s'. It's index: %d\n", wordForSearch, isExist)
	}
	fmt.Printf("Word '%s' is %d in input\n", wordForCount, countOccurrences)
	fmt.Printf("Input after replate '%s' for **:\n%v", wordForReplace, replacedInput)

}
func toTitleCase(s string) string {
	words := strings.Fields(strings.ToLower(s))
	for i := range words {
		if len(words[i]) > 0 {
			runes := []rune(words[i])
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}
