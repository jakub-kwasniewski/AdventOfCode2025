package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Move struct {
	Direction string
	Steps     int
}

func ReadFileLines(filePath string) ([]Move, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error during file opening: ", err)
		return nil, err
	}
	defer file.Close()

	var lines []Move

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error during number conversion:", err)
			return nil, err
		}
		lines = append(lines, Move{Direction: direction, Steps: steps})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("File reading error:", err)
		return nil, err
	}

	return lines, nil
}

func partOne(moves []Move) int {
	initialPosition := 50
	points := 0
	currentPosition := initialPosition
	for _, move := range moves {
		if move.Direction == "L" {
			move.Steps *= -1
		}
		currentPosition = (currentPosition + move.Steps + 100) % 100
		if currentPosition == 0 {
			points += 1
		}
	}
	return points
}

func partTwo(moves []Move) int {
	initialPosition := 50
	points := 0
	currentPosition := initialPosition
	for _, move := range moves {
		points += move.Steps / 100
		newSteps := move.Steps % 100
		if move.Direction == "L" {
			newSteps *= -1
		}
		newPosition := currentPosition + newSteps
		if newPosition <= 0 || newPosition >= 100 {
			if currentPosition != 0 {
				points += 1
			}
		}
		currentPosition = (currentPosition + newSteps + 100) % 100
	}
	return points
}

func main() {
	lines, err := ReadFileLines("input.txt")
	if err != nil {
		return
	}
	fmt.Println("Part 1 solution: ", partOne(lines))
	fmt.Println("Part 2 solution: ", partTwo(lines))
}
