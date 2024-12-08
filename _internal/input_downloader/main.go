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

func writeFile(content []byte, year, day int, fileName string) {
	parsedDay := strconv.Itoa(day)
	if len(parsedDay) < 2 {
		parsedDay = "0" + strconv.Itoa(day)
	}

	inputFilePath := fmt.Sprintf("./%d/Day-%s/%s", year, parsedDay, fileName)
	resolvedInputFilePath, err := filepath.Abs(inputFilePath)

	err = os.MkdirAll(filepath.Dir(resolvedInputFilePath), 0755)
	if err != nil {
		panic(err)
	}

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

func putBasicFiles(year, day int) {
	// Write go.mod file
	dayS := strconv.Itoa(day)
	if len(dayS) < 2 {
		dayS = "0" + dayS
	}

	goModContent := fmt.Sprintf(`module github.com/LordMartron94/Advent-of-Code/%d/Day-%s

go 1.23
`, year, dayS)

	writeFile([]byte(goModContent), year, day, "go.mod")

	rawContent := `package main

import (
	"fmt"
	"os"

	"github.com/LordMartron94/Advent-of-Code/${year}/Day-${day}/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/${year}/Day-${day}/pipeline/pipes"
	"github.com/LordMartron94/Advent-of-Code/${year}/Day-${day}/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	pipeline2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/patterns/pipeline"
)

const year = ${year}
const day = ${dayS}

func main() {
	utilities.ChangeWorkingDirectoryToSpecificTask(year, day)
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

	pipesToRun := []pipeline2.Pipe[common.PipelineContext[task_rules.LexingTokenType]]{
		&pipes.GetInputDataPipe{},
	}

	startingContext := common.NewPipelineContext[task_rules.LexingTokenType](file)
	pipeline := pipeline2.NewPipeline(pipesToRun)
	result := pipeline.Process(*startingContext)

	fmt.Println(fmt.Sprintf("Final result (task 1): %d", result.Result))
}
`
	dayPaddedString := strconv.Itoa(day)

	goMainContent := strings.Replace(rawContent, "${dayS}", dayPaddedString, -1)

	if len(dayPaddedString) < 2 {
		dayPaddedString = "0" + dayPaddedString
	}

	goMainContent = strings.Replace(goMainContent, "${year}", strconv.Itoa(year), -1)
	goMainContent = strings.Replace(goMainContent, "${day}", dayPaddedString, -1)

	writeFile([]byte(goMainContent), year, day, "main.go")
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

	writeFile(response, year, day, "input.txt")
	putBasicFiles(year, day)
}
