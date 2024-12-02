package main

import (
	"fmt"
	"os"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-01/pipeline/common"
	pipes2 "github.com/LordMartron94/Advent-of-Code/2024/Day-01/pipeline/pipes"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	pipeline2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/patterns/pipeline"
)

const year = 2024
const day = 1

func main() {
	const expectedSumDistance = 2164381
	const expectedSumIncrease = 20719933
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

	pipes := []pipeline2.Pipe[common.PipelineContext]{
		&pipes2.GetInputDataPipe{},
		&pipes2.TransformSlicesPipe{},
		&pipes2.CalculationPipe{},
	}

	startingContext := common.NewPipelineContext(file)

	pipeline := pipeline2.NewPipeline(pipes)
	finalData := pipeline.Process(*startingContext)

	fmt.Printf("Total distance for the tokens: %d\n", finalData.TotalDistance)
	fmt.Printf("Sum of increases: %d\n", finalData.TotalIncreases)

	if finalData.TotalDistance != expectedSumDistance {
		fmt.Println("Incorrect sum of distances")
	}

	if finalData.TotalIncreases != expectedSumIncrease {
		fmt.Println("Incorrect sum of increases")
	}
}
