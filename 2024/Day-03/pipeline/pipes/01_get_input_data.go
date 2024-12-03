package pipes

import (
	"log"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-03/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-03/task_rules"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/default_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
)

type GetInputDataPipe struct {
}

// Process method to process the input. Input is the filepath.
func (g *GetInputDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	lexingRules := []default_rules.LexingRuleInterface[task_rules.LexingTokenType]{
		&task_rules.MulKeywordRuleLex{},
		&task_rules.DontKeywordRuleLex{},
		&task_rules.DoKeywordRuleLex{},
		&task_rules.OpenParenthesisRuleLex{},
		&task_rules.CloseParenthesisRuleLex{},
		&task_rules.CommaRuleLex{},
		&task_rules.DigitRuleLex{},
		&task_rules.InvalidTokenLex{},
	}

	parsingRules := []rules.ParsingRuleInterface[task_rules.LexingTokenType]{
		&task_rules.MultiplyOperationRuleParser{},
		&task_rules.DontRuleParser{},
		&task_rules.DoRuleParser{},
		&task_rules.InvalidTokenRuleParser{},
	}

	fileHandler := utilities.NewFileHandler[task_rules.LexingTokenType](input.Reader, lexingRules, parsingRules, task_rules.EOFToken)
	//tokens := fileHandler.Lex()
	tree, err := fileHandler.Parse()

	//for i, token := range tokens {
	//	fmt.Println(fmt.Sprintf("Token (%d)... Type: %d, value: '%s'", i+1, token.Type, token.Value))
	//}

	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	//tree.Print(2)

	input.ParseTree = tree
	return input
}
