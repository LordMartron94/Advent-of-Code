package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// isValidYear checks if a string is a valid year.
// A valid year is a 4-digit number.
func isValidYear(year string) bool {
	// Check if the string is 4 digits long.
	if len(year) != 4 {
		return false
	}

	// Check if the string contains only digits.
	if !regexp.MustCompile(`^\d+$`).MatchString(year) {
		return false
	}

	// Check if the year is within a reasonable range.
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	if yearInt < 1000 || yearInt > 9999 {
		return false
	}

	return true
}

func getYearDirs() ([]string, error) {
	dirEntries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	yearDirs := make([]string, 0)

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			if isValidYear(dirEntry.Name()) {
				yearDirs = append(yearDirs, dirEntry.Name())
			}
		}
	}

	return yearDirs, nil
}

func isValidDay(day string) bool {
	// Check if the string is 2 digits long.
	if len(day) != 2 {
		return false
	}

	// Check if the string contains only digits.
	if !regexp.MustCompile(`^\d+$`).MatchString(day) {
		return false
	}

	// Check if the day is within a reasonable range.
	dayInt, err := strconv.Atoi(day)
	if err != nil {
		return false
	}
	if dayInt < 1 || dayInt > 31 {
		return false
	}

	return true
}

func getDayDirs(yearDir string) ([]string, error) {
	dirEntries, err := os.ReadDir(yearDir)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	dayDirs := make([]string, 0)

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			parts := strings.Split(dirEntry.Name(), "-")
			numDay := parts[1]
			// 0-pad numDay
			if len(numDay) < 2 {
				numDay = "0" + numDay
			}

			if parts[0] != "Day" || !isValidDay(numDay) {
				continue
			}

			dayDirs = append(dayDirs, dirEntry.Name())
		}
	}

	return dayDirs, nil
}

type EntryPoint struct {
	Day   int
	Path  string
	Label string
}

func getTasksForYear(yearDir string) []EntryPoint {
	dayDirs, err := getDayDirs(yearDir)

	entryPointsList := make([]EntryPoint, 0)

	if err != nil {
		log.Printf("Error reading day directories for year: %s\n", yearDir)
		return entryPointsList
	}

	yearNumber := strings.TrimPrefix(yearDir, "./")
	year, _ := strconv.Atoi(yearNumber)

	for _, dayDir := range dayDirs {
		// Get day number from directory name
		dayNum := strings.TrimPrefix(dayDir, "Day-")
		day, _ := strconv.Atoi(dayNum)

		entryPoint, err, label := getEntryPointForDay(filepath.Join(yearDir, dayDir), year, day)

		if err != nil {
			log.Printf("Error reading entry point for day: %s\n", dayDir)
			continue
		}

		// Append to the slice
		entryPointsList = append(entryPointsList, EntryPoint{
			Day:   day,
			Path:  entryPoint,
			Label: label,
		})
	}

	sort.Slice(entryPointsList, func(i, j int) bool {
		return entryPointsList[i].Day < entryPointsList[j].Day
	})

	return entryPointsList
}

func getEntryPointForDay(dayDir string, year int, day int) (string, error, string) {
	// Find main.go file in the day directory
	mainGoFilePath := filepath.Join(dayDir, "main.go")
	resolvedMainGoFilePath, err := filepath.Abs(mainGoFilePath)

	if err != nil {
		log.Printf("Error resolving main.go file path: %s\n", dayDir)
		log.Println(err)
		return "", err, ""
	}

	if _, err := os.Stat(resolvedMainGoFilePath); err == nil {
		return resolvedMainGoFilePath, nil, fmt.Sprintf("Year '%d' - Day '%d'", year, day)
	} else if os.IsNotExist(err) {
		fmt.Printf("No main.go file found in day directory: %s\n", dayDir)
	} else {
		log.Printf("Error reading main.go file in day directory: %s\n", dayDir)
		log.Println(err)
	}

	return "", nil, ""
}

func processTasks(tasks []EntryPoint) {
	for i, task := range tasks {
		fmt.Printf("Task %d): %s\n", i+1, task.Label)
	}

	fmt.Println("------------------------------")

	fmt.Println("Total tasks:", len(tasks))

	// Let user select a task to execute
	var selectedTaskIndex int
	reader := bufio.NewReader(os.Stdin)
	maxNum := strconv.Itoa(len(tasks))
	fmt.Print("Enter the number of the task you want to execute (1-" + maxNum + ") [" + maxNum + "]: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input) // Remove newline and potential leading/trailing spaces

	selectedTaskIndex = len(tasks)
	if input != "" {
		selectedTaskIndex, err = strconv.Atoi(input)
		if err != nil {
			log.Println("Invalid input:", err)
			return
		}
	}

	if selectedTaskIndex < 1 || selectedTaskIndex > len(tasks) {
		fmt.Println("Invalid task number")
		return
	}

	selectedTaskExe := tasks[selectedTaskIndex-1].Path

	cmd := exec.Command("go", "run", selectedTaskExe)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Printf("Error executing task: %s\n", selectedTaskExe)
		log.Println(err)
		return
	}
}

func main() {
	yearDirs, err := getYearDirs()

	if err != nil {
		log.Fatal(err)
		return
	}

	tasks := make([]EntryPoint, 0)

	for _, yearDir := range yearDirs {
		tasksForYear := getTasksForYear(yearDir)
		tasks = append(tasks, tasksForYear...)
	}

	processTasks(tasks)
}
