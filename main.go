package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Concatenate strings
	var letters = []string{"s", "u", "b", "c", "r", "i", "b", "e"}

	// Test time to concatenate strings
	duration, word := timeToConcat(letters)
	fmt.Println("It took", duration, "to concatenate the word:", word)

	// Test time using strings.Builder
	buildDuration, buildWord := timeUsingStringBuilder(letters)
	fmt.Println("It took", buildDuration, "to build the word:", buildWord)

	// Test time to join letters
	joinDuration, joinWord := timeUsingJoin(letters)
	fmt.Println("It took", joinDuration, "to join the word:", joinWord)

}

// Concatenate strings with loop
func timeToConcat(letters []string) (duration time.Duration, outputString string) {
	t0 := time.Now()
	word := ""
	for i := range letters {
		word += letters[i]
	}
	return time.Since(t0), word
}

// Concatenate strings using strings.Builder
func timeUsingStringBuilder(letters []string) (duration time.Duration, outputString string) {
	t0 := time.Now()
	var stringBuilder strings.Builder
	for i := range letters {
		stringBuilder.WriteString(letters[i])
	}
	word := stringBuilder.String()
	return time.Since(t0), word
}

// Concatenate strings using strings.Join
func timeUsingJoin(letters []string) (duration time.Duration, outputString string) {
	t0 := time.Now()
	word := strings.Join(letters, "")
	return time.Since(t0), word
}
