package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/LordMartron94/Advent-of-Code/_internal/input_downloader/env_parser"
	requester2 "github.com/LordMartron94/Advent-of-Code/_internal/input_downloader/requester"
)

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(fmt.Sprintf("%s: ", prompt))
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error reading input:", err)
		os.Exit(1)
	}

	userInput = strings.TrimSpace(userInput)

	return userInput
}

func getEnvFilePath(args []string) string {
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <env_file_path>")

		path := getUserInput("Enter the path to the environment file")
		if path == "" {
			fmt.Println("Invalid path")
			os.Exit(1)
		}

		return path
	}

	return args[1]
}

func cleanFilePath(envFilePath string) {
	envFilePath = filepath.Clean(envFilePath)

	if _, err := os.Stat(envFilePath); os.IsNotExist(err) {
		fmt.Printf("File not found: %s\n", envFilePath)
		os.Exit(1)
	}
}

func getDayAndYear() (int, int) {
	now := time.Now()
	finalYear := now.Year()
	finalDay := now.Day()

	chosenYear := getUserInput(fmt.Sprintf("Enter the year (e.g., 2024, leave empty for default) [%d]", finalYear))
	chosenDay := getUserInput(fmt.Sprintf("Enter the day (e.g., 15, leave empty for default) [%d]", finalDay))

	if chosenYear != "" {
		converted, err := strconv.Atoi(chosenYear)

		if err != nil {
			fmt.Println("Error parsing year:", err)
			os.Exit(1)
		}

		finalYear = converted
	}

	if chosenDay != "" {
		converted, err := strconv.Atoi(chosenDay)

		if err != nil {
			fmt.Println("Error parsing day:", err)
			os.Exit(1)
		}

		finalDay = converted
	}

	// Ensure year is not in the future and day is between 1 and 25
	if finalYear > time.Now().Year() || (finalYear == time.Now().Year() && finalDay > 25) {
		fmt.Println("Invalid year or day")
		os.Exit(1)
	}

	return finalYear, finalDay
}

func writeInputFile(content []byte, year, day int) {
	inputFilePath := fmt.Sprintf("./%d/Day %d/input.txt", year, day)
	resolvedInputFilePath, err := filepath.Abs(inputFilePath)

	inputFile, err := os.OpenFile(resolvedInputFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		panic(err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Println("Error closing the file:", err)
			os.Exit(1)
		}
	}(inputFile)

	_, err = inputFile.Write(content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Input file created successfully: ", inputFilePath)
}

func main() {
	args := os.Args
	envFilePath := getEnvFilePath(args)
	cleanFilePath(envFilePath)

	envFile, err := os.OpenFile(envFilePath, os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer func(envFile *os.File) {
		err := envFile.Close()
		if err != nil {
			fmt.Println("Error closing the file:", err)
			os.Exit(1)
		}
	}(envFile)

	keyValuePairs := env_parser.GetKeyValuePairs(envFile)

	sessionToken := keyValuePairs["SESSION_TOKEN"]
	requester := requester2.Requester{SessionToken: &sessionToken}

	year, day := getDayAndYear()

	response, err := requester.Get(day, year)

	if err != nil {
		panic(err)
	}

	writeInputFile(response, year, day)
}
