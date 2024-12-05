package pipes

import (
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-05/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-05/task_rules"
	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/common_transformers"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformDataPipe struct {
}

func (t *TransformDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	pairs := make([][]shared3.Token[task_rules.LexingTokenType], 0)
	updateCommands := make([][]shared3.Token[task_rules.LexingTokenType], 0)

	callbackFinder := func(node *shared.ParseTree[task_rules.LexingTokenType]) (shared2.TransformCallback[task_rules.LexingTokenType], int) {
		switch node.Symbol {
		case "pair_of_numbers":
			return common_transformers.GetPairsFromSpecificChildren("first_number", "second_number", &pairs), 0
		case "update_command":
			return common_transformers.CollectRowByChildSymbols([]string{"number"}, &updateCommands), 0
		}
		return nil, 0
	}

	transformer := transforming.NewTransformer(
		callbackFinder,
	)
	transformer.Transform(input.ParseTree)

	//for _, pair := range pairs {
	//	fmt.Printf("Pair: %s, %s\n", pair[0].Value, pair[1].Value)
	//}
	//fmt.Println("=================")
	//for _, updateCommand := range updateCommands {
	//	for _, token := range updateCommand {
	//		fmt.Printf("UpdateContents command: %s\n", token.Value)
	//	}
	//	fmt.Println("----------------")
	//}

	transformedPairs := make([][]int, 0)
	transformedUpdateCommands := make([][]int, 0)

	for _, pair := range pairs {
		firstNumber, err := strconv.Atoi(string(pair[0].Value))
		if err != nil {
			panic(err)
		}
		secondNumber, err := strconv.Atoi(string(pair[1].Value))
		if err != nil {
			panic(err)
		}
		transformedPairs = append(transformedPairs, []int{firstNumber, secondNumber})
	}

	for _, updateCommand := range updateCommands {
		command := make([]int, 0)

		for _, token := range updateCommand {
			commandNumber, err := strconv.Atoi(string(token.Value))

			if err != nil {
				panic(err)
			}

			command = append(command, commandNumber)
		}

		transformedUpdateCommands = append(transformedUpdateCommands, command)
	}

	input.Manuals = transformedPairs
	input.Updates = transformedUpdateCommands

	return input
}
