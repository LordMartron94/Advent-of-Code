package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/default_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/common_callbacks"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

const year = "2024"
const day = "Day-01"

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

func getNumAppearancesInSlice(slice []int, target int) int {
	result := 0

	for _, num := range slice {
		if num == target {
			result++
		}
	}

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

func GetSlicesFromParseTree(tree shared.ParseTree) ([]int, []int) {
	num1Slice := make([]int, 0)
	num2Slice := make([]int, 0)

	callbackFinder := func(node *shared.ParseTree) shared2.TransformCallback {
		switch node.Symbol {
		case "first_number":
			return common_callbacks.AppendTokenValueToSlice(&num1Slice, strconv.Atoi)
		case "second_number":
			return common_callbacks.AppendTokenValueToSlice(&num2Slice, strconv.Atoi)
		}
		return nil
	}

	transformer := transforming.NewTransformer(callbackFinder)
	transformer.Transform(&tree)

	return num1Slice, num2Slice
}

func main() {
	err := os.Chdir(fmt.Sprintf("./%s/%s", year, day))

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

	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0644)
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

	lexingRules := make([]default_rules.LexingRuleInterface, 0)
	lexingRules = append(lexingRules, &default_rules.WhitespaceRule{})
	lexingRules = append(lexingRules, &default_rules.DigitRule{})

	parsingRules := []rules.ParsingRuleInterface{
		&rules.PairRule{},
		&rules.WhitespaceRule{},
		&rules.NumberRule{},
	}

	fileHandler := utilities.NewFileHandler(file, lexingRules, parsingRules)

	parseTree, err := fileHandler.Parse()

	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	num1Slice, num2Slice := GetSlicesFromParseTree(*parseTree)

	// Sort both lists in ascending order
	sort.Ints(num1Slice)
	sort.Ints(num2Slice)

	distances := transformPairsToDistances(num1Slice, num2Slice)

	totalDistance := sum(distances)

	fmt.Printf("Total distance for the tokens: %d\n", totalDistance)

	appearancesMap := getAppearancesMap(num1Slice, num2Slice)
	increases := getIncreases(appearancesMap, num1Slice)
	sumIncreases := sum(increases)

	fmt.Printf("Sum of increases: %d\n", sumIncreases)
}

func getIncreases(appearancesMap map[int]int, slice []int) []int {
	increases := make([]int, len(slice))

	for i, num := range slice {
		increases[i] = num * appearancesMap[num]
	}

	return increases
}
