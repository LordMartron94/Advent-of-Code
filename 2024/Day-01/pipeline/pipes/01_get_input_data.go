package pipes

import (
	"log"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-01/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
)

type GetInputDataPipe struct {
}

// Process method to process the input. Input is the filepath.
func (g *GetInputDataPipe) Process(input common.PipelineContext) common.PipelineContext {
	lexingRules := make([]rules.LexingRuleInterface, 0)
	lexingRules = append(lexingRules, &rules.WhitespaceRule{})
	lexingRules = append(lexingRules, &rules.DigitRule{})

	parsingRules := []rules.ParsingRuleInterface{
		&rules.PairRule{},
		&rules.WhitespaceRule{},
		&rules.NumberRule{},
	}

	fileHandler := utilities.NewFileHandler(input.Reader, lexingRules, parsingRules)
	tree, err := fileHandler.Parse()

	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	input.ParseTree = tree
	return input
}
