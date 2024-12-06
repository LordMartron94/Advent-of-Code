package pipes

import (
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type CalculateDataPipe struct {
}

type Position struct {
	rIndex int
	cIndex int
	deltaC int
	deltaR int
}

type Grid struct {
	rows [][]shared.Token[task_rules.LexingTokenType]
}

func (c *CalculateDataPipe) GetIndicesOfStartingPosition(grid Grid) Position {
	for r, row := range grid.rows {
		for c, token := range row {
			if token.Type == task_rules.CarrotToken {
				return Position{rIndex: r, cIndex: c, deltaC: 0, deltaR: -1}
			}
		}
	}

	panic("No starting position found")
}

func (c *CalculateDataPipe) IsLoop(currentPos Position, path []Position) bool {
	// Admittedly, I got stuck here and got Rodos from ArjanCodes to give a hint.
	// Thank you Rodos, you're a hero!
	// I do understand the logic behind it. Because if you face the same direction from the same position,
	// with these rules, you will guaranteed follow the same path again, thus leading to the same output... A loop.

	if slices.Contains(path, currentPos) {
		return true
	}

	return false
}

func (c *CalculateDataPipe) Seen(currentPos Position, path []Position) bool {
	for _, pos := range path {
		if pos.rIndex == currentPos.rIndex && pos.cIndex == currentPos.cIndex {
			return true
		}
	}

	return false
}

func (c *CalculateDataPipe) Solve(pos Position, visitedNodes *int, grid Grid) bool {
	rowCount := len(grid.rows) - 1
	columnCount := len(grid.rows[0]) - 1
	path := make([]Position, 0)

	for {
		if !c.Seen(pos, path) {
			path = append(path, pos)
			*visitedNodes++
		}

		// Bounds Check
		if pos.rIndex+pos.deltaR > rowCount || pos.rIndex+pos.deltaR < 0 ||
			pos.cIndex+pos.deltaC > columnCount || pos.cIndex+pos.deltaC < 0 {
			return false
		}
		if grid.rows[pos.rIndex+pos.deltaR][pos.cIndex+pos.deltaC].Type == task_rules.HashToken {
			pos.deltaC, pos.deltaR = -pos.deltaR, pos.deltaC
		} else {
			pos.rIndex += pos.deltaR
			pos.cIndex += pos.deltaC
		}
		if c.IsLoop(pos, path) {
			return true
		}
	}
}

func (c *CalculateDataPipe) GetAmountOfLoopsForAlteredGrids(startingPos Position, originalGrid Grid, numOfGoroutineWorkers int) int {
	numOfLoopPotentials := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	jobs := make(chan job, len(originalGrid.rows)*len(originalGrid.rows[0]))

	// Start workers
	for i := 0; i < numOfGoroutineWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				numMoves := 0
				looping := c.Solve(j.startingPos, &numMoves, j.grid)
				if looping {
					mu.Lock()
					numOfLoopPotentials++
					mu.Unlock()
				}
			}
		}()
	}

	// Generate jobs
	for i, row := range originalGrid.rows {
		for j, token := range row {
			if token.Type == task_rules.DotToken {
				gridCopy := copyGrid(originalGrid)
				gridCopy.rows[i][j].Type = task_rules.HashToken
				jobs <- job{startingPos: startingPos, grid: gridCopy}
			}
		}
	}
	close(jobs)

	wg.Wait()
	return numOfLoopPotentials
}

type job struct {
	startingPos Position
	grid        Grid
}

// Helper function to copy the grid
func copyGrid(grid Grid) Grid {
	newGrid := Grid{rows: make([][]shared.Token[task_rules.LexingTokenType], len(grid.rows))}
	for i, row := range grid.rows {
		newGrid.rows[i] = make([]shared.Token[task_rules.LexingTokenType], len(row))
		copy(newGrid.rows[i], row)
	}
	return newGrid
}

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	const originalTotalExecutionTimeInMS = 289459

	grid := Grid{rows: input.Rows}

	startingPos := c.GetIndicesOfStartingPosition(grid)

	var numMoves int

	startTime := time.Now()
	_ = c.Solve(startingPos, &numMoves, grid)

	input.Result = numMoves
	input.BlockResult = c.GetAmountOfLoopsForAlteredGrids(startingPos, grid, 18)

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution time: %vµs\n", executionTime.Microseconds())

	optimizationMS := (float64(originalTotalExecutionTimeInMS)/float64(executionTime.Milliseconds()) - 1) * 100
	fmt.Printf("Optimization improvement: %.2f%%\n", optimizationMS)

	return input
}
