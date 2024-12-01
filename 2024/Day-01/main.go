package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
)

const year = "2024"
const day = "Day-01"

func readInputFile() []string {
	// Set current working directory
	err := os.Chdir(fmt.Sprintf("./%s/%s", year, day))

	// Print current working directory
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		os.Exit(1)
	}

	fmt.Println("Current working directory:", dir)

	if err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}

	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0644)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			os.Exit(1)
		}
	}(f)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseInput(lines []string) ([]int, []int) {
	num1Slice := make([]int, len(lines))
	num2Slice := make([]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, "   ")
		for j, part := range parts {
			if j == 0 {
				num1, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("Error parsing number:", err)
					os.Exit(1)
				}
				num1Slice[i] = num1
			} else if j == 1 {
				num2, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("Error parsing number:", err)
					os.Exit(1)
				}
				num2Slice[i] = num2
			} else {
				panic("Invalid input format")
			}
		}
	}

	return num1Slice, num2Slice
}

func transformPairsToDistances(num1Slice, num2Slice []int) []int {
	distances := make([]int, len(num1Slice))

	for i := 0; i < len(num1Slice); i++ {
		// Calculate absolute difference between two numbers
		dist := num1Slice[i] - num2Slice[i]

		if dist < 0 {
			dist = -dist
		}

		distances = append(distances, dist)
	}

	return distances
}

func sum(distances []int) int {
	sum := 0
	for _, distance := range distances {
		sum += distance
	}
	return sum
}

func sliceContains(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}

func getNumAppearancesInSlice(slice []int, target int) int {
	result := 0

	for _, num := range slice {
		if num == target {
			result++
		}
	}

	//fmt.Println(fmt.Sprintf("Gotten appearances for target %d: %d", target, result))
	return result
}

func mapContainsKey(m map[int]int, key int) bool {
	_, ok := m[key]

	if ok {
		return true
	}
	return false
}

func getAppearancesMap(num1Slice, num2Slice []int) map[int]int {
	appearances := make(map[int]int)

	for _, num := range num1Slice {
		if !mapContainsKey(appearances, num) {
			appearances[num] = getNumAppearancesInSlice(num2Slice, num)
		} else {
			continue
		}
	}

	return appearances
}

func main() {
	// Set current working directory
	err := os.Chdir(fmt.Sprintf("./%s/%s", year, day))

	// Print current working directory
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		os.Exit(1)
	}

	fmt.Println("Current working directory:", dir)

	if err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}

	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			os.Exit(1)
		}
	}(file)

	fileHandler := utilities.NewFileHandler(file)

	tokens := fileHandler.Lex()

	for _, token := range tokens {
		fmt.Println(fmt.Sprintf("Token (%d) - %s", token.Type, token.Value))
	}

	//const numOneResult = 2226302
	//
	//lines := readInputFile()
	//num1Slice, num2Slice := parseInput(lines)
	//
	//// Sort both lists in ascending order
	//sort.Ints(num1Slice)
	//sort.Ints(num2Slice)
	//
	//distances := transformPairsToDistances(num1Slice, num2Slice)
	//
	//totalDistance := sum(distances)
	//
	//fmt.Printf("Total distance for the tokens: %d\n", totalDistance)
	//
	//appearancesMap := getAppearancesMap(num1Slice, num2Slice)
	//increases := getIncreases(appearancesMap, num1Slice)
	//sumIncreases := sum(increases)
	//
	//fmt.Printf("Sum of increases: %d\n", sumIncreases)
	//
	//if sumIncreases == numOneResult {
	//	fmt.Println("Solution for test 1 is correct.")
	//} else {
	//	fmt.Println("Solution for test 1 is incorrect.")
	//}
}

func getIncreases(appearancesMap map[int]int, slice []int) []int {
	increases := make([]int, len(slice))

	for i, num := range slice {
		increases[i] = num * appearancesMap[num]
	}

	return increases
}
