package pipes

import (
	"fmt"
	"time"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type CalculateDataPipe struct {
}

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	const originalTotalExecutionTimeInMS = 289459

	// TODO - Create factory for pathfinding rules
	pathFreeFunc := func(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[shared.Token[task_rules.LexingTokenType]]) bool {
		pos := finder.GetPositionInDirectory(currentPosition, currentDirection, 1)

		if finder.OutOfBounds(pos) {
			return true
		}

		item := finder.GetItemAtPosition(pos)

		if item.Type != task_rules.HashToken {
			return true
		}

		return false
	}

	getNewDirectionFunc := func(_ matrix.Position, currentDirection pathfinding.Direction) pathfinding.Direction {
		return currentDirection.Turn()
	}

	ruleset := pathfinding.Ruleset[shared.Token[task_rules.LexingTokenType]]{
		IsBasic: true,
		Rules: []pathfinding.Rule[shared.Token[task_rules.LexingTokenType]]{
			{
				MatchFunc: func(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[shared.Token[task_rules.LexingTokenType]]) bool {
					return pathFreeFunc(currentPosition, currentDirection, finder)
				},
				GetNewPosition: func(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[shared.Token[task_rules.LexingTokenType]]) matrix.Position {
					return finder.GetPositionInDirectory(currentPosition, currentDirection, 1)
				},
				GetNewDirection: func(_ matrix.Position, currentDirection pathfinding.Direction) pathfinding.Direction {
					return currentDirection
				},
			},
			{
				MatchFunc: func(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[shared.Token[task_rules.LexingTokenType]]) bool {
					return !pathFreeFunc(currentPosition, currentDirection, finder)
				},
				GetNewPosition: func(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[shared.Token[task_rules.LexingTokenType]]) matrix.Position {
					return finder.GetPositionInDirectory(currentPosition, getNewDirectionFunc(currentPosition, currentDirection), 1)
				},
				GetNewDirection: func(currentPosition matrix.Position, currentDirection pathfinding.Direction) pathfinding.Direction {
					return getNewDirectionFunc(currentPosition, currentDirection)
				},
			},
		},
	}

	pathFinderHelper := pathfinding.NewPathFinder(input.Rows, shared.Token[task_rules.LexingTokenType].Equals, ruleset)

	startTime := time.Now()
	startToken := shared.Token[task_rules.LexingTokenType]{Type: task_rules.CarrotToken, Value: []byte("^")}
	dotToken := shared.Token[task_rules.LexingTokenType]{Type: task_rules.DotToken, Value: []byte(".")}
	hashToken := shared.Token[task_rules.LexingTokenType]{Type: task_rules.HashToken, Value: []byte("#")}

	numMoves, err := pathFinderHelper.GetNumberOfUniqueNodesVisitedUntilOutOfBounds(startToken, pathfinding.Up)

	if err != nil {
		panic(err)
	}

	numOfVariationsLooping, err := pathFinderHelper.GetNumberOfLoopingMatricesForGeneratedVariations(startToken, pathfinding.Up, dotToken, hashToken)

	if err != nil {
		panic(err)
	}

	input.Result = numMoves
	input.BlockResult = numOfVariationsLooping

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution time: %vµs\n", executionTime.Microseconds())

	optimizationMS := (float64(originalTotalExecutionTimeInMS)/float64(executionTime.Milliseconds()) - 1) * 100
	fmt.Printf("Optimization improvement: %.2f%%\n", optimizationMS)

	return input
}
