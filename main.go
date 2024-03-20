package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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

func getDigits(line string) int {
	var firstDigit string
	var lastDigit string

	firstDigitChanel := make(chan string)
	go getFirstDigit(firstDigitChanel, line)

	firstDigit = <-firstDigitChanel

	lastDigitChanel := make(chan string)
	go getLastDigit(lastDigitChanel, line)

	lastDigit = <-lastDigitChanel

	digits, _ := strconv.Atoi(firstDigit + lastDigit)

	return digits
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//totalChanel receives the digit from each line from the workers
	totalChanel := make(chan int)

	//lines chanel send each line of the file to the workers,
	lines := make(chan string, 100)

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)

	start := time.Now()

	//sending the lines to the workers through the lines chanel
	go func() {
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()

	workNumber := 100

	for i := 0; i < workNumber; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for line := range lines {
				lineDigits := getDigits(line)
				totalChanel <- lineDigits
			}
		}()
	}

	go func() {
		wg.Wait()
		close(totalChanel)
	}()

	totalSum := 0
	for sum := range totalChanel {
		totalSum += sum
	}

	fmt.Println("Total sum:", totalSum)

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}
