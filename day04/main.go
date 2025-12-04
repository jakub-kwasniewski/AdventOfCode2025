package main

import (
	"bufio"
	"fmt"
	"os"
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

func partOne(lines []string) int {
	rowsNum := len(lines)
	columnsNum := len(lines[0])
	iter := []int{-1, 0, 1}
	answer := 0
	for r, line := range lines {
		for c, char := range line {
			if char == '@' {
				totalRools := 0
				for _, dr := range iter {
					if r+dr < 0 || r+dr >= rowsNum {
						continue
					}
					for _, dc := range iter {
						if c+dc < 0 || c+dc >= columnsNum {
							continue
						}
						if dc == 0 && dr == 0 {
							continue
						}
						if lines[r+dr][c+dc] == '@' {
							totalRools += 1
						}
					}
				}
				if totalRools < 4 {
					answer += 1
				}
			}
		}
	}
	return answer
}

func partTwo(banks []string) int {
	return 0
}

func main() {
	lines, err := ReadFileLines("input.txt")
	if err != nil {
		return
	}
	fmt.Println("Part 1 solution: ", partOne(lines))
	fmt.Println("Part 2 solution: ", partTwo(lines))
}
