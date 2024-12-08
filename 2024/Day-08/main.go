package main

import (
	"fmt"
	"os"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-08/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-08/pipeline/pipes"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-08/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	pipeline2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/patterns/pipeline"
)

const year = 2024
const day = 8

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

	pipesToRun := []pipeline2.Pipe[common.PipelineContext[task_rules.LexingTokenType]]{
		&pipes.GetInputDataPipe{},
		&pipes.TransformDataPipe{},
		&pipes.CalculateDataPipe{},
	}

	startingContext := common.NewPipelineContext[task_rules.LexingTokenType](file)
	pipeline := pipeline2.NewPipeline(pipesToRun)
	result := pipeline.Process(*startingContext)

	fmt.Println(fmt.Sprintf("Final result (task 1): %d", result.Result))
	fmt.Println(fmt.Sprintf("Final result (task 2): %d", result.ResultPart2))
}
