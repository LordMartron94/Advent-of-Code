package pipes

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-06/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/common_transformers"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformDataPipe struct{}

func (t *TransformDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	rows := make([][]shared.Token[task_rules.LexingTokenType], 0)

	callbackFinder := func(node *shared3.ParseTree[task_rules.LexingTokenType]) (shared2.TransformCallback[task_rules.LexingTokenType], int) {
		switch node.Symbol {
		case "row":
			return common_transformers.AppendChildrenToSlice(&rows), 0
		}
		return nil, 0
	}

	transformer := transforming.NewTransformer(
		callbackFinder,
	)
	transformer.Transform(input.ParseTree)

	numDots := 0
	for _, row := range rows {
		for _, token := range row {
			if token.Type == task_rules.DotToken {
				numDots++
			}
		}
	}

	fmt.Printf("Number of visible dots: %d\n", numDots)

	input.Rows = rows

	return input
}
