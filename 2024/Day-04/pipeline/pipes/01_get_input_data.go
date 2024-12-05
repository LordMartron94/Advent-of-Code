package pipes

import (
	"log"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-04/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-04/task_rules"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	rules2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
)

type GetInputDataPipe struct {
}

// Process method to process the input. Input is the filepath.
func (g *GetInputDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	lexerRuleFactory := task_rules.NewRuleset()
	parsingRuleFactory := task_rules.NewParsingRuleFactory()

	lexingRules := []rules.LexingRuleInterface[task_rules.LexingTokenType]{
		lexerRuleFactory.GetXCharRuleLex(),
		lexerRuleFactory.GetMCharRuleLex(),
		lexerRuleFactory.GetACharRuleLex(),
		lexerRuleFactory.GetSCharRuleLex(),
		lexerRuleFactory.GetNewLineRuleLexer(),
		lexerRuleFactory.GetInvalidTokenRuleLex(),
	}

	parsingRules := []rules2.ParsingRuleInterface[task_rules.LexingTokenType]{
		parsingRuleFactory.GetHorizontalLineParserRule(),
		parsingRuleFactory.GetNewLineTokenParserRule(),
		parsingRuleFactory.GetInvalidTokenRule(),
	}

	fileHandler := utilities.NewFileHandler[task_rules.LexingTokenType](input.Reader, lexingRules, parsingRules, task_rules.IgnoreToken)
	//tokens, err := fileHandler.Lex()
	////
	//if err != nil {
	//	log.Fatalf("error lexing input: %v", err)
	//}
	//
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

	tree, err := fileHandler.Parse()

	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	//tree.Print(2)

	//os.Exit(0)

	input.ParseTree = tree
	return input
}
