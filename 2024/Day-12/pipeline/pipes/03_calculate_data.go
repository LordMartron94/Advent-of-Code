package pipes

import (
	"github.com/LordMartron94/Advent-of-Code/2024/Day-12/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-12/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type CalculateDataPipe struct {
}

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	lazyComparer := func(a, b shared.Token[task_rules.LexingTokenType]) bool {
		return a.Type == b.Type
	}
	strictComparer := func(a, b shared.Token[task_rules.LexingTokenType]) bool {
		return a.Equals(b)
	}

	matrixHelper := matrix.NewMatrixHelper(input.Rows, lazyComparer)
	outerPlots := make([]int, 0)
	regions := matrixHelper.GetRegions(&strictComparer, &outerPlots)

	cost := 0
	discountCost := 0

	for i, region := range regions {
		regionArea := len(region)
		perimeter := outerPlots[i]
		edgeCount := matrixHelper.GetNumberOfEdgesAroundPolygon(region)

		cost += regionArea * perimeter
		discountCost += regionArea * edgeCount
	}

	input.Result = cost
	input.Result2 = discountCost

	return input
}
