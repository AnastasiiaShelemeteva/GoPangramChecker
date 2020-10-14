package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var alphabet [26]string = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {

	fmt.Println("Please, write the sentence, you want to check: ")

	checkPangram(missingLetters(matchLetters(takeInput())))
}

func takeInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentenceToCheck := scanner.Text()
	sentenceToCheck = strings.ToLower(sentenceToCheck)

	return sentenceToCheck
}

func matchLetters(input string) []string {

	arrayToCheck := strings.Split(input, "")

	var matchLetters []string

	for i := 0; i < len(alphabet); i++ {
		for j := 0; j < len(arrayToCheck); j++ {
			if alphabet[i] == arrayToCheck[j] {
				matchLetters = append(matchLetters, alphabet[i])
				break
			} else {
				continue
			}

		}

	}

	return matchLetters

}

func missingLetters(matchLetters []string) []string {
	var missingLetters []string
	for _, value := range alphabet {
		for j, value2 := range matchLetters {
			if value != value2 {
				missingLetters = append(missingLetters, value)
				break

			} else {
				matchLetters = append(matchLetters[:j], matchLetters[j+1:]...)
				break
			}
		}

	}

	return missingLetters
}

func checkPangram(missingLetters []string) {
	if len(missingLetters) == 0 {
		fmt.Println("This is a pangram!")
	} else {
		fmt.Printf("This is not a pangram. Missing letters are %v", strings.Join(missingLetters, ", "))
	}
}
