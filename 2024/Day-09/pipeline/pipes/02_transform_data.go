package pipes

import (
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-09/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-09/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/common_transformers"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformDataPipe struct{}

func (t *TransformDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	pairs := make([][]shared.Token[task_rules.LexingTokenType], 0)

	callbackFinder := func(node *shared3.ParseTree[task_rules.LexingTokenType]) (shared2.TransformCallback[task_rules.LexingTokenType], int) {
		switch node.Symbol {
		case "row":
			return common_transformers.GetPairsFromChildren("pair", &pairs), 0
		}
		return nil, 0
	}

	transformer := transforming.NewTransformer(
		callbackFinder,
	)
	transformer.Transform(input.ParseTree)

	elements := make([]*int, 0)

	for i, pair := range pairs {
		fileSize, _ := strconv.Atoi(string(pair[0].Value))
		freeBlockSize, _ := strconv.Atoi(string(pair[1].Value))

		for range fileSize {
			elements = append(elements, &i)
		}

		for range freeBlockSize {
			elements = append(elements, nil)
		}
	}

	input.DiskInfo = elements

	return input
}
