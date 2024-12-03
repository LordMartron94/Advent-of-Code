package pipes

import (
	"github.com/LordMartron94/Advent-of-Code/2024/Day-03/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-03/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/common_calculations"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/common_transformers"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type CalculateDataPipe struct {
}

func (t *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	enabled := true

	multiplications := make([]int, 0)
	enabledMultiplications := make([]int, 0)

	callbackFinder := func(node *shared.ParseTree[task_rules.LexingTokenType]) (shared2.TransformCallback[task_rules.LexingTokenType], int) {
		switch node.Symbol {
		case "multiply_operation":
			return common_transformers.CombineCallbacks[task_rules.LexingTokenType](
				common_transformers.ApplyBinaryOperationToChildren[task_rules.LexingTokenType](func(left, right int) int {
					return left * right
				}, &multiplications),
				common_transformers.ApplyBinaryOperationToChildren[task_rules.LexingTokenType](func(left, right int) int {
					if enabled {
						return left * right
					}
					return 0
				}, &enabledMultiplications)), 0
		case "do":
			return func(_ *shared.ParseTree[task_rules.LexingTokenType]) {
				enabled = true
			}, 0
		case "don't":
			return func(_ *shared.ParseTree[task_rules.LexingTokenType]) {
				enabled = false
			}, 0
		}
		return nil, 0
	}

	transformer := transforming.NewTransformer(
		callbackFinder,
	)
	transformer.Transform(input.ParseTree)

	var total int
	common_calculations.SumInts(&multiplications, &total)

	var totalBool int
	common_calculations.SumInts(&enabledMultiplications, &totalBool)

	input.Result = total
	input.EnabledResult = totalBool
	return input
}
