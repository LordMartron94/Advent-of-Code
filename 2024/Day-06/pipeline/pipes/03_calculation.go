package pipes

import (
	"fmt"
	"time"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/factory"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type CalculateDataPipe struct {
}

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	const originalTotalExecutionTimeInMS = 289459

	startToken := shared.Token[task_rules.LexingTokenType]{Type: task_rules.CarrotToken, Value: []byte("^")}
	dotToken := shared.Token[task_rules.LexingTokenType]{Type: task_rules.DotToken, Value: []byte(".")}
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

	pathFinderHelper := pathfinding.NewPathFinder(input.Rows, shared.Token[task_rules.LexingTokenType].Equals, ruleset, false)

	startTime := time.Now()

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
