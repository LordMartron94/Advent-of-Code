package pipes

import (
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-10/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-10/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/graph"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type CalculateDataPipe struct{}

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	lazyComparer := func(a, b shared.Token[task_rules.LexingTokenType]) bool {
		return a.Type == b.Type
	}
	strictComparer := func(a, b shared.Token[task_rules.LexingTokenType]) bool {
		return a.Equals(b)
	}

	startNode := shared.Token[task_rules.LexingTokenType]{Type: task_rules.NumberToken, Value: []byte("0")}
	endNode := shared.Token[task_rules.LexingTokenType]{Type: task_rules.NumberToken, Value: []byte("9")}

	graphHelper := graph.NewGraphHelper(input.Rows, lazyComparer)
	validPathsUnique := graphHelper.FindSuitablePathsBetweenNodes(startNode, endNode, func(start, end shared.Token[task_rules.LexingTokenType]) bool {
		sV, _ := strconv.Atoi(string(start.Value))
		eV, _ := strconv.Atoi(string(end.Value))

		diff := eV - sV
		return diff == 1
	}, strictComparer, true)

	validPathsAll := graphHelper.FindSuitablePathsBetweenNodes(startNode, endNode, func(start, end shared.Token[task_rules.LexingTokenType]) bool {
		sV, _ := strconv.Atoi(string(start.Value))
		eV, _ := strconv.Atoi(string(end.Value))

		diff := eV - sV
		return diff == 1
	}, strictComparer, false)

	input.Result = len(validPathsUnique)
	input.Result2 = len(validPathsAll)

	return input
}
