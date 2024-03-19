package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// used in the first part of the code challenge
// func getDigits(line string) (first string, last string) {
// 	for _, char := range line {
// 		if unicode.IsDigit(char) {
// 			first = string(char)
// 			break
// 		}
// 	}

// 	for i := len(line) - 1; i >= 0; i-- {
// 		if unicode.IsDigit(rune(line[i])) {
// 			last = string(line[i])
// 			break
// 		}
// 	}

// 	return
// }

func getDigits(line string) (first string, last string) {
	validDigits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	firstDigitIndex := -1
	lastDigitIndex := -1

	for index, char := range line {
		if unicode.IsDigit(char) {
			first = string(char)
			firstDigitIndex = index
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			last = string(line[i])
			lastDigitIndex = i
			break
		}
	}

	for index, value := range validDigits {
		foundIndex := strings.Index(line, index)
		if foundIndex >= 0 {
			if foundIndex < firstDigitIndex {
				firstDigitIndex = foundIndex
				first = strconv.Itoa(value)
			}
		}

		foundIndex = strings.LastIndex(line, index)
		if foundIndex >= 0 {
			if foundIndex > lastDigitIndex {
				lastDigitIndex = foundIndex
				last = strconv.Itoa(value)
			}
		}
	}

	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, secondDigit := getDigits(line)
		number, err := strconv.Atoi(firstDigit + secondDigit)
		if err != nil {
			panic(err)
		}

		total += number
	}

	fmt.Printf("The total amout is %d\n", total)
}
