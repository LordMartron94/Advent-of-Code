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
type Parser[T comparable] struct {
	lexer    *lexing.Lexer[T]
	ruleSet  *rules.Ruleset[T]
	stateMap map[rules.ParsingRuleInterface[T]]fsm.State[ParsingStateArgs[T]]
}

// NewParser creates a new parser from the given input
func NewParser[T comparable](lexer *lexing.Lexer[T], parsingRules []rules.ParsingRuleInterface[T]) *Parser[T] {
	parser := &Parser[T]{
		lexer:   lexer,
		ruleSet: rules.NewRuleset[T](parsingRules),
	}

	stateMap, err := parser.generateFSM()
	if err != nil {
		panic(fmt.Sprintf("failed to generate FSM: %v", err))
	}

	parser.stateMap = stateMap

	return parser
}

func startState[T comparable](ctx context.Context, args ParsingStateArgs[T]) (ParsingStateArgs[T], fsm.State[ParsingStateArgs[T]], error) {
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
func (p *Parser[T]) Parse() (*shared2.ParseTree[T], error) {
	// Reset lexer to be sure it works
	p.lexer.Reset()
	tokens := p.lexer.GetTokens()
	fmt.Println("Starting Parsing Process...")

	args := ParsingStateArgs[T]{
		tokens:       tokens,
		currentToken: nil,
		currentIndex: 0,
		currentBuffer: &shared2.ParseTree[T]{
			Symbol: "root",
		},
		parser: p,
	}

	args, err := fsm.Run(context.Background(), args, startState)
	if err != nil {
		return nil, fmt.Errorf("parsing failed: %w", err)
	}

	fmt.Println("Parsing Process Complete!")

	return args.currentBuffer, nil
}

// ParsingStateArgs holds the arguments for the parsing FSM
type ParsingStateArgs[T comparable] struct {
	parser        *Parser[T]
	tokens        []*shared.Token[T]
	currentToken  *shared.Token[T]
	currentIndex  int
	currentBuffer *shared2.ParseTree[T]
}

// generateFSM generates the FSM for parsing
func (p *Parser[T]) generateFSM() (map[rules.ParsingRuleInterface[T]]fsm.State[ParsingStateArgs[T]], error) {
	stateMap := make(map[rules.ParsingRuleInterface[T]]fsm.State[ParsingStateArgs[T]])

	for _, rule := range p.ruleSet.Rules {
		stateMap[rule] = func(ctx context.Context, args ParsingStateArgs[T]) (ParsingStateArgs[T], fsm.State[ParsingStateArgs[T]], error) {
			if args.currentIndex >= len(args.tokens) {
				return args, nil, nil
			}

			args.currentToken = args.tokens[args.currentIndex]

			node, err, consumed := rule.Match(args.tokens, args.currentIndex)
			if err != nil {
				return args, nil, fmt.Errorf("rule %s failed to match: %w", rule.Symbol(), err)
			}
			if node == nil {
				args.currentIndex += consumed
				return args, startState, nil
			}

			args.currentBuffer.Children = append(args.currentBuffer.Children, node)
			args.currentIndex += consumed

			return args, startState, nil
		}
	}

	return stateMap, nil
}
