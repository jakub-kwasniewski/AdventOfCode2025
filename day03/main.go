package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error during file opening: ", err)
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("File reading error:", err)
		return nil, err
	}

	return lines, nil
}

func partOne(banks []string) int {
	totalJoltage := 0
	battQuant := len(banks[0])
	for i, bank := range banks {
		fmt.Print(i, ":")
		firstDigit, _ := strconv.Atoi(string(bank[0]))
		secondDigit, _ := strconv.Atoi(string(bank[battQuant-1]))
		for i := 1; i < battQuant-1; i++ {
			currentDigit, _ := strconv.Atoi(string(bank[i]))
			if currentDigit > firstDigit {
				firstDigit = currentDigit
				secondDigit, _ = strconv.Atoi(string(bank[battQuant-1]))
			} else if currentDigit > secondDigit {
				secondDigit = currentDigit
			}
		}
		fmt.Println(firstDigit*10 + secondDigit)
		totalJoltage += firstDigit*10 + secondDigit
	}
	return totalJoltage
}

func partTwo(banks []string) int {
	totalJoltage := 0
	battQuant := len(banks[0])
	for _, bank := range banks {
		startIdx := 0
		bankJoltage := 0
		for i := 11; i >= 0; i-- {
			hiJoltage := 0
			for j := startIdx; j < battQuant-i; j++ {
				current, _ := strconv.Atoi(string(bank[j]))
				if current > hiJoltage {
					hiJoltage = current
					startIdx = j + 1
				}
			}
			bankJoltage += hiJoltage * int(math.Pow(10, float64(i)))
		}
		totalJoltage += bankJoltage
	}
	return totalJoltage
}

func main() {
	lines, err := ReadFileLines("input.txt")
	if err != nil {
		return
	}
	fmt.Println("Part 1 solution: ", partOne(lines))
	fmt.Println("Part 2 solution: ", partTwo(lines))
}
