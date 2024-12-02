package pipes

import (
	"sort"
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-01/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/common_calculations"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/common_transformers"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformSlicesPipe struct {
}

func calculateDistance(num1, num2 int) int {
	d := num1 - num2

	if d < 0 {
		d = -d
	}

	return d
}

func (t *TransformSlicesPipe) Process(input common.PipelineContext) common.PipelineContext {
	num1Slice := make([]int, 0)
	num2Slice := make([]int, 0)
	distances := make([]int, 0)

	callbackFinder := func(node *shared.ParseTree) (shared2.TransformCallback, int) {
		switch node.Symbol {
		case "first_number":
			return common_transformers.AppendTokenValueToSliceSorted(&num1Slice, strconv.Atoi, sort.Ints), 0
		case "second_number":
			return common_transformers.AppendTokenValueToSliceSorted(&num2Slice, strconv.Atoi, sort.Ints), 0
		}
		return nil, 0
	}

	transformer := transforming.NewTransformer(
		callbackFinder,
	)
	transformer.Transform(input.ParseTree)

	pairs := common_calculations.GetPairs(num1Slice, num2Slice)

	for _, pair := range pairs {
		distances = append(distances, calculateDistance(pair[0], pair[1]))
	}

	input.Column1Slice = num1Slice
	input.Column2Slice = num2Slice
	input.Distances = distances

	return input
}
