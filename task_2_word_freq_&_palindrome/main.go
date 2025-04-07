package main

import (
	"fmt"
	"strings"
)

func main() {

	var userOptions int

	userOptions = selectingOption(userOptions)

	switch userOptions {
	case 1:
		var word string
		fmt.Print("Enter a word to check if it's a palindrome: ")
		fmt.Scanln(&word)
		if isPalindrome(word) {
			fmt.Printf("The word '%s' is a palindrome.\n", word)
		} else {
			fmt.Printf("The word '%s' is not a palindrome.\n", word)
		}
	case 2:
		var sentence string
		fmt.Print("Enter a sentence to count word frequency: ")
		fmt.Scanln(&sentence)
		wordFrequencyMap := wordFrequency(sentence)
		fmt.Println("Word Frequency:")
		for word, count := range wordFrequencyMap {
			fmt.Printf("'%s': %d\n", word, count)
		}
	case 3:
		fmt.Println("Exiting the program. Goodbye!")
		return
	default:
		fmt.Println("Invalid option. Please select a valid option (1-3).")

	}

}

func isPalindrome(word string) bool {
	var reverseWord = reverseWordFunc(word)

	if word == reverseWord {
		return true
	}

	return false

}

func wordFrequency(word string) map[string]int {

	var wordFrequencyDictionary = make(map[string]int)

	var processedWord strings.Builder

	lowercaseWord := strings.ToLower(word)

	// removing punctuation ,spaces
	// var processedWord string

	for _, char := range lowercaseWord {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == ' ' {
			processedWord.WriteRune(char)
		}
	}

	words := strings.Fields(processedWord.String())
	for _, w := range words {
		wordFrequencyDictionary[w]++
	}

	return wordFrequencyDictionary
}

func reverseWordFunc(word string) string {
	var wordArray []string = strings.Split(word, "")
	var reverseWord string

	for i := len(wordArray) - 1; i >= 0; i-- {
		reverseWord += wordArray[i]
	}

	return reverseWord
}

func selectingOption(userOptions int) int {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Welcome to the word frequency and palindrome checker!")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Please select an option:")
	fmt.Println("1. Check if a word is a palindrome")
	fmt.Println("2. Count the frequency of words in a sentence")
	fmt.Println("3. Exit")
	fmt.Println("-------------------------------------------------------")
	fmt.Print("Enter your choice (1-3): ")
	fmt.Scanln(&userOptions)
	fmt.Println("-------------------------------------------------------")

	return userOptions
}
