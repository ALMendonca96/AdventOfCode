package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var validDigits = map[string]int{
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

func getFirstDigit(ch chan string, line string) {
	var digit string
	var digitIndex int
	for index, char := range line {
		if unicode.IsDigit(char) {
			digit = string(char)
			digitIndex = index
			break
		}
	}

	for index, value := range validDigits {
		foundIndex := strings.Index(line, index)
		if foundIndex >= 0 {
			if foundIndex < digitIndex {
				digitIndex = foundIndex
				digit = strconv.Itoa(value)
			}
		}
	}

	ch <- digit
}

func getLastDigit(ch chan string, line string) {
	var digit string
	var digitIndex int
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			digit = string(line[i])
			digitIndex = i
			break
		}
	}

	for index, value := range validDigits {
		foundIndex := strings.LastIndex(line, index)
		if foundIndex >= 0 {
			if foundIndex > digitIndex {
				digitIndex = foundIndex
				digit = strconv.Itoa(value)
			}
		}
	}

	ch <- digit
}

func getDigitsAsync(ch chan int, line string) {
	var firstDigit string
	var lastDigit string

	firstDigitChanel := make(chan string)
	go getFirstDigit(firstDigitChanel, line)

	firstDigit = <-firstDigitChanel

	lastDigitChanel := make(chan string)
	go getLastDigit(lastDigitChanel, line)

	lastDigit = <-lastDigitChanel

	digits, _ := strconv.Atoi(firstDigit + lastDigit)
	ch <- digits
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int
	var rows int
	for scanner.Scan() {
		line := scanner.Text()

		chanel := make(chan int)
		go getDigitsAsync(chanel, line)
		total += <-chanel
		rows++
	}

	fmt.Printf("The total amout is %d\n", total)
}
