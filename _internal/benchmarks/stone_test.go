package benchmarks

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-11/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-11/pipeline/pipes"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-11/task_rules"
	pipeline2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/patterns/pipeline"
)

func BenchmarkStoneSuite(b *testing.B) {
	tmpFile, err := os.CreateTemp("", "benchmark-stone")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	originalStdout := os.Stdout

	os.Stdout = tmpFile

	prepared := prepareStone()

	b.Run("Solve 20x", func(b *testing.B) {
		benchmarkStone(b, 20, prepared, pipes.CalculateDataPipe{})
	})

	//b.Run("Solve 40x", func(b *testing.B) {
	//	benchmarkStone(b, 40, prepared)
	//})

	os.Stdout = originalStdout
}

func prepareStone() common.PipelineContext[task_rules.LexingTokenType] {
	fileDir, _ := os.Getwd()
	currentDir := filepath.Dir(fileDir)
	rootDir := filepath.Join(currentDir, "..")
	filePath := filepath.Join(rootDir, "2024", "Day-11", "input.txt")

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
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
	}

	startingContext := common.NewPipelineContext[task_rules.LexingTokenType](file)
	pipeline := pipeline2.NewPipeline(pipesToRun)
	return pipeline.Process(*startingContext)
}

func benchmarkStone(b *testing.B, iterations int, input common.PipelineContext[task_rules.LexingTokenType], pipe pipes.CalculateDataPipe) {
	shutdownSignal := make(chan struct{})
	wg := &sync.WaitGroup{}

	b.ResetTimer() // Start timer after initialization
	for i := 0; i < b.N; i++ {
		pipe.GetStoneCount(input, iterations)
	}

	close(shutdownSignal)
	wg.Wait()
}
