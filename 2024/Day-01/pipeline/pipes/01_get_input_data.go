package pipes

import (
	"log"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-01/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-01/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	rules2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
)

type GetInputDataPipe struct {
}

// Process method to process the input. Input is the filepath.
func (g *GetInputDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	lexingRuleFactory := task_rules.NewRuleFactory()
	parsingRuleFactory := task_rules.NewParsingRuleFactory()

	lexingRules := []rules.LexingRuleInterface[task_rules.LexingTokenType]{
		lexingRuleFactory.GetWhitespaceRuleLex(),
		lexingRuleFactory.GetDigitRuleLex(),
		lexingRuleFactory.GetInvalidTokenRuleLex(),
	}

	parsingRules := []rules2.ParsingRuleInterface[task_rules.LexingTokenType]{
		parsingRuleFactory.GetPairParsingRule(),
		parsingRuleFactory.GetInvalidTokenParsingRule(),
	}

	fileHandler := utilities.NewFileHandler(input.Reader, lexingRules, parsingRules, task_rules.IgnoreToken)
	tree, err := fileHandler.Parse()

	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	//tree.Print(2)
	//os.Exit(0)

	input.ParseTree = tree
	return input
}
