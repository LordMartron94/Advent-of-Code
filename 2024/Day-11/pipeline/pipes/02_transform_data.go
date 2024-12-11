package pipes

import (
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-11/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-11/task_rules"
	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/common_transformers"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformDataPipe struct {
}

func (t *TransformDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	stones := make([]int, 0)

	callbackFinder := func(node *shared3.ParseTree[task_rules.LexingTokenType]) (shared2.TransformCallback[task_rules.LexingTokenType], int) {
		switch node.Symbol {
		case "number":
			return common_transformers.AppendTokenValueToSlice[int, task_rules.LexingTokenType](&stones, strconv.Atoi), 0
		}
		return nil, 0
	}

	transformer := transforming.NewTransformer(
		callbackFinder,
	)
	transformer.Transform(input.ParseTree)

	input.Stones = stones

	return input
}
