package parsing

import (
	"context"
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/patterns/fsm"
)

// Parser is a struct to represent a parser
type Parser struct {
	lexer    *lexing.Lexer
	ruleSet  *rules.Ruleset
	stateMap map[rules.ParsingRuleInterface]fsm.State[ParsingStateArgs]
}

// NewParser creates a new parser from the given input
func NewParser(lexer *lexing.Lexer, parsingRules []rules.ParsingRuleInterface) *Parser {
	parser := &Parser{
		lexer:   lexer,
		ruleSet: rules.NewRuleset(parsingRules),
	}

	stateMap, err := parser.generateFSM()
	if err != nil {
		panic(fmt.Sprintf("failed to generate FSM: %v", err))
	}

	parser.stateMap = stateMap

	return parser
}

func startState(ctx context.Context, args ParsingStateArgs) (ParsingStateArgs, fsm.State[ParsingStateArgs], error) {
	if args.currentIndex >= len(args.tokens) {
		return args, nil, nil
	}

	rule, err := args.parser.ruleSet.GetMatchingRule(args.tokens, args.currentIndex)

	if err != nil {
		return args, nil, fmt.Errorf("no matching rule found: %w", err)
	}

	fn := args.parser.stateMap[rule]

	return fn(ctx, args)
}

// Parse parses the input and returns the parse tree
func (p *Parser) Parse() (*shared2.ParseTree, error) {
	// Reset lexer to be sure it works
	p.lexer.Reset()
	tokens := p.lexer.GetTokens()

	args := ParsingStateArgs{
		tokens:       tokens,
		currentToken: nil,
		currentIndex: 0,
		currentBuffer: &shared2.ParseTree{
			Symbol: "root",
		},
		parser: p,
	}

	args, err := fsm.Run(context.Background(), args, startState)
	if err != nil {
		return nil, fmt.Errorf("parsing failed: %w", err)
	}

	return args.currentBuffer, nil
}

// ParsingStateArgs holds the arguments for the parsing FSM
type ParsingStateArgs struct {
	parser        *Parser
	tokens        []*shared.Token
	currentToken  *shared.Token
	currentIndex  int
	currentBuffer *shared2.ParseTree
}

// generateFSM generates the FSM for parsing
func (p *Parser) generateFSM() (map[rules.ParsingRuleInterface]fsm.State[ParsingStateArgs], error) {
	stateMap := make(map[rules.ParsingRuleInterface]fsm.State[ParsingStateArgs])

	for _, rule := range p.ruleSet.Rules {
		stateMap[rule] = func(ctx context.Context, args ParsingStateArgs) (ParsingStateArgs, fsm.State[ParsingStateArgs], error) {
			if args.currentIndex >= len(args.tokens) {
				return args, nil, nil
			}

			args.currentToken = args.tokens[args.currentIndex]

			node, err := rule.Match(args.tokens, args.currentIndex)
			if err != nil {
				return args, nil, fmt.Errorf("rule %s failed to match: %w", rule.Symbol(), err)
			}
			if node == nil {
				args.currentIndex += 1
				return args, startState, nil
			}

			args.currentBuffer.Children = append(args.currentBuffer.Children, node)
			args.currentIndex += len(node.Children) + 1

			return args, startState, nil
		}
	}

	return stateMap, nil
}
