package pipes

import (
	"log"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-03/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-03/task_rules"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	rules2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
)

type GetInputDataPipe struct {
}

// Process method to process the input. Input is the filepath.
func (g *GetInputDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	lexerRuleFactory := task_rules.NewRuleset()

	lexingRules := []rules.LexingRuleInterface[task_rules.LexingTokenType]{
		lexerRuleFactory.GetMulKeywordRuleLex(),
		lexerRuleFactory.GetDontKeywordRuleLex(),
		lexerRuleFactory.GetDoKeywordRuleLex(),
		lexerRuleFactory.GetOpenParenthesisRuleLex(),
		lexerRuleFactory.GetCloseParenthesisRuleLex(),
		lexerRuleFactory.GetCommaRuleLex(),
		lexerRuleFactory.GetDigitRuleLex(),
		lexerRuleFactory.GetInvalidTokenRuleLex(),
	}

	parsingRules := []rules2.ParsingRuleInterface[task_rules.LexingTokenType]{
		&task_rules.MultiplyOperationRuleParser{},
		&task_rules.DontRuleParser{},
		&task_rules.DoRuleParser{},
		&task_rules.InvalidTokenRuleParser{},
	}

	fileHandler := utilities.NewFileHandler[task_rules.LexingTokenType](input.Reader, lexingRules, parsingRules)
	//tokens, err := fileHandler.Lex()

	//if err != nil {
	//	log.Fatalf("error lexing input: %v", err)
	//}

	//for i, token := range tokens {
	//	fmt.Println(fmt.Sprintf("Token (%d)... Type: %d, value: '%s'", i+1, token.Type, token.Value))
	//}

	tree, err := fileHandler.Parse()

	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	//tree.Print(2)

	input.ParseTree = tree
	return input
}
