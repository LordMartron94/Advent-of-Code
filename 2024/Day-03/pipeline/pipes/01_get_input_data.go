package pipes

import (
	"fmt"
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
	lexingRules := []rules.LexingRuleInterface[task_rules.LexingTokenType]{
		//&task_rules.MulKeywordRuleLex{},
		//&task_rules.DontKeywordRuleLex{},
		//&task_rules.DoKeywordRuleLex{},
		//&task_rules.OpenParenthesisRuleLex{},
		//&task_rules.CloseParenthesisRuleLex{},
		//&task_rules.CommaRuleLex{},
		//&task_rules.DigitRuleLex{},
		&task_rules.InvalidTokenLex{},
	}

	parsingRules := []rules2.ParsingRuleInterface[task_rules.LexingTokenType]{
		&task_rules.MultiplyOperationRuleParser{},
		&task_rules.DontRuleParser{},
		&task_rules.DoRuleParser{},
		&task_rules.InvalidTokenRuleParser{},
	}

	fileHandler := utilities.NewFileHandler[task_rules.LexingTokenType](input.Reader, lexingRules, parsingRules)
	tokens, err := fileHandler.Lex()

	if err != nil {
		log.Fatalf("error lexing input: %v", err)
	}

	for i, token := range tokens {
		fmt.Println(fmt.Sprintf("Token (%d)... Type: %d, value: '%s'", i+1, token.Type, token.Value))
	}

	if string(tokens[len(tokens)-1].Value) != "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))" {
		fmt.Println("Test failed")
	}

	//tree, err := fileHandler.Parse()

	//if err != nil {
	//	log.Fatalf("error parsing input: %v", err)
	//}

	//tree.Print(2)

	//input.ParseTree = tree
	return input
}
