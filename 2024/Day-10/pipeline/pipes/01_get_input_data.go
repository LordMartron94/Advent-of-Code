package pipes

import (
	"log"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-10/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-10/task_rules"
	rules2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
)

type GetInputDataPipe struct {
}

// Process method to process the input. Input is the filepath.
func (g *GetInputDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	lexerRuleFactory := task_rules.NewRuleset()
	parsingRuleFactory := task_rules.NewParsingRuleFactory()

	lexingRules := []rules.LexingRuleInterface[task_rules.LexingTokenType]{
		lexerRuleFactory.GetCharTokenRuleLex(),
		lexerRuleFactory.GetNewLineTokenRuleLex(),
		lexerRuleFactory.GetInvalidTokenRuleLex(),
	}

	parsingRules := []rules2.ParsingRuleInterface[task_rules.LexingTokenType]{
		parsingRuleFactory.GetRowParsingRule(),
		parsingRuleFactory.GetInvalidTokenRule(),
	}

	fileHandler := utilities.NewFileHandler[task_rules.LexingTokenType](input.Reader, lexingRules, parsingRules, task_rules.IgnoreToken)
	//tokens, err := fileHandler.Lex()
	//
	//if err != nil {
	//	log.Fatalf("error lexing input: %v", err)
	//}

	//for i, token := range tokens {
	//	if token.Type == task_rules.NewLineToken {
	//		fmt.Println("Encountered New Line Token")
	//		continue
	//	} else if token.Type == task_rules.IgnoreToken {
	//		fmt.Println("Encountered Ignore Token")
	//		continue
	//	}
	//
	//	fmt.Println(fmt.Sprintf("Token (%d)... Type: %d, value: '%s'", i+1, token.Type, token.Value))
	//}
	//
	//os.Exit(0)

	tree, err := fileHandler.Parse()

	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	//tree.Print(2, []task_rules.LexingTokenType{task_rules.NewLineToken})
	////
	//os.Exit(0)

	input.ParseTree = tree
	return input
}
