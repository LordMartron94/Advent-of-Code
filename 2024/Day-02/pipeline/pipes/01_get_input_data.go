package pipes

import (
	"log"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-02/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-02/task_rules"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/default_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
)

type GetInputDataPipe struct {
}

// Process method to process the input. Input is the filepath.
func (g *GetInputDataPipe) Process(input common.PipelineContext) common.PipelineContext {
	lexingRules := []default_rules.LexingRuleInterface{
		&task_rules.NewlineRuleLex{},
		&task_rules.SpaceRuleLex{},
		&default_rules.DigitRule{},
	}

	parsingRules := []rules.ParsingRuleInterface{
		&task_rules.NewLineRuleParser{},
		&task_rules.SpaceRuleParser{},
		&task_rules.ReportRule{},
	}

	fileHandler := utilities.NewFileHandler(input.Reader, lexingRules, parsingRules)
	tree, err := fileHandler.Parse()

	//tree.Print(2)

	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	input.ParseTree = tree
	return input
}
