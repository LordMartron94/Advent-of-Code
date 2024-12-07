package benchmarks

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/pipeline/pipes"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/factory"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	pipeline2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/patterns/pipeline"
)

func BenchmarkLoggingSuite(b *testing.B) {
	tmpFile, err := os.CreateTemp("", "benchmark-solve")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	originalStdout := os.Stdout

	os.Stdout = tmpFile

	prepared := prepare()
	pathFinder := getPathFinder(prepared)

	startToken := shared.Token[task_rules.LexingTokenType]{Type: task_rules.CarrotToken, Value: []byte("^")}
	startDirection := pathfinding.Up

	b.Run("Solve 1x", func(b *testing.B) {
		benchmarkSolve(b, 1, pathFinder, startToken, startDirection)
	})

	b.Run("Solve 100x", func(b *testing.B) {
		benchmarkSolve(b, 100, pathFinder, startToken, startDirection)
	})

	b.Run("Solve 10000x", func(b *testing.B) {
		benchmarkSolve(b, 10000, pathFinder, startToken, startDirection)
	})

	os.Stdout = originalStdout
}

func prepare() common.PipelineContext[task_rules.LexingTokenType] {
	fileDir, _ := os.Getwd()
	currentDir := filepath.Dir(fileDir)
	rootDir := filepath.Join(currentDir, "..")
	filePath := filepath.Join(rootDir, "2024", "Day-06", "test.txt")

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

func getPathFinder(pipelineContext common.PipelineContext[task_rules.LexingTokenType]) *pathfinding.PathFinder[shared.Token[task_rules.LexingTokenType]] {
	//startToken := shared.Token[task_rules.LexingTokenType]{Type: task_rules.CarrotToken, Value: []byte("^")}
	//dotToken := shared.Token[task_rules.LexingTokenType]{Type: task_rules.DotToken, Value: []byte(".")}
	hashToken := shared.Token[task_rules.LexingTokenType]{Type: task_rules.HashToken, Value: []byte("#")}

	ruleFactory := factory.NewPathfindingRuleFactory[shared.Token[task_rules.LexingTokenType]]()

	pathFreeFunc := func(finder pathfinding.PathFinder[shared.Token[task_rules.LexingTokenType]], nextTile shared.Token[task_rules.LexingTokenType]) bool {
		return !finder.EqualityCheck(nextTile, hashToken)
	}

	ruleset := pathfinding.PathfindingRuleset[shared.Token[task_rules.LexingTokenType]]{
		IsBasic: true,
		Rules: []pathfinding.PathfindingRuleInterface[shared.Token[task_rules.LexingTokenType]]{
			ruleFactory.GetBasicRule(pathFreeFunc, func(currentDirection pathfinding.Direction) pathfinding.Direction {
				return currentDirection
			}, 1),
			ruleFactory.GetBasicRule(func(finder pathfinding.PathFinder[shared.Token[task_rules.LexingTokenType]], nextTile shared.Token[task_rules.LexingTokenType]) bool {
				return !pathFreeFunc(finder, nextTile)
			}, func(currentDirection pathfinding.Direction) pathfinding.Direction {
				return currentDirection.TurnRight()
			}, 1),
		},
	}

	return pathfinding.NewPathFinder(pipelineContext.Rows, shared.Token[task_rules.LexingTokenType].Equals, ruleset, false)
}

func benchmarkSolve(b *testing.B, iterations int, pathFinder *pathfinding.PathFinder[shared.Token[task_rules.LexingTokenType]], startToken shared.Token[task_rules.LexingTokenType], startDirection pathfinding.DirectionExternal) {
	shutdownSignal := make(chan struct{})
	wg := &sync.WaitGroup{}

	b.ResetTimer() // Start timer after initialization
	for i := 0; i < b.N; i++ {
		for j := 0; j < iterations; j++ {
			_, err := pathFinder.GetNumberOfUniqueNodesVisitedUntilOutOfBounds(startToken, startDirection)
			if err != nil {
				continue
			}
		}
	}

	close(shutdownSignal)
	wg.Wait()
}
