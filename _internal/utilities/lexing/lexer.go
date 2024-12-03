package lexing

import (
	"bufio"
	"bytes"
	"context"
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/default_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/patterns/fsm"
)

// TODO - Refactor for readability and maintainability
// TODO - IMPORTANT - add capacity for easy multicharacter lexing... Really difficult now.
// Take inspiration from the parser - it's really easy there

type LexerStateArgs[T comparable] struct {
	lexer        *Lexer[T]
	eofTokenType T

	CurrentToken  *shared.Token[T]
	currentBuffer *bytes.Buffer
}

// Lexer is a struct for lexing input.
type Lexer[T comparable] struct {
	runes        []rune
	index        int
	eofTokenType T

	ruleSet  *default_rules.Ruleset[T]
	stateMap map[default_rules.LexingRuleInterface[T]]fsm.State[LexerStateArgs[T]]
}

// NewLexer creates a new Lexer with the given reader.
func NewLexer[T comparable](reader io.Reader, rules []default_rules.LexingRuleInterface[T], eofTokenType T) *Lexer[T] {
	bReader := bufio.NewReader(reader)

	runes := make([]rune, 0)

	for {
		scannedRune, _, err := bReader.ReadRune()

		if err == io.EOF {
			break
		}

		runes = append(runes, scannedRune)
	}

	lexer := &Lexer[T]{
		runes:        runes,
		index:        -1,
		eofTokenType: eofTokenType,
	}

	lexer.ruleSet = default_rules.NewRuleset[T](rules)

	sM, err := lexer.generateFSM()

	if err != nil {
		panic(err)
	}

	lexer.stateMap = sM

	return lexer
}

func (l *Lexer[T]) GetRuleset() default_rules.Ruleset[T] {
	return *l.ruleSet
}

func (l *Lexer[T]) generateFSM() (map[default_rules.LexingRuleInterface[T]]fsm.State[LexerStateArgs[T]], error) {
	stateMap := make(map[default_rules.LexingRuleInterface[T]]fsm.State[LexerStateArgs[T]])

	for _, rule := range l.ruleSet.Rules {
		stateMap[rule] = func(ctx context.Context, args LexerStateArgs[T]) (LexerStateArgs[T], fsm.State[LexerStateArgs[T]], error) {
			cRune, err := args.lexer.Consume()
			args, _, err = handleConsumeErr(err, args)

			if err != nil {
				if err == io.EOF {
					t := rule.CreateToken(args.currentBuffer)
					args.CurrentToken = t
					return args, nil, err
				}

				return args, nil, err
			}

			matchedRule2, err := args.lexer.ruleSet.GetMatchingRule(cRune, l)
			if err != nil {
				return args, nil, err
			}

			if matchedRule2 == rule {
				args.currentBuffer.WriteRune(cRune)
				return args, l.stateMap[matchedRule2], nil
			} else {
				args.lexer.Pushback()
				t := rule.CreateToken(args.currentBuffer)
				args.CurrentToken = t
				return args, nil, nil
			}
		}
	}

	return stateMap, nil
}

// Peek returns the next rune without advancing the lexer's index.
func (l *Lexer[T]) Peek() (rune, error) {
	if l.index+1 >= len(l.runes) {
		return ' ', io.EOF
	}

	return l.runes[l.index+1], nil
}

// PeekN returns the next N runes without advancing the lexer's index.
func (l *Lexer[T]) PeekN(n int) ([]rune, error) {
	if l.index+n >= len(l.runes) {
		return nil, io.EOF
	}

	return l.runes[l.index+1 : l.index+n+1], nil
}

// Consume returns the next rune and advances the lexer's index.
func (l *Lexer[T]) Consume() (rune, error) {
	if l.index+1 >= len(l.runes) {
		return ' ', io.EOF
	}

	l.index++

	return l.runes[l.index], nil
}

// Pushback reverts the lexer's index to the previous position.
func (l *Lexer[T]) Pushback() {
	if l.index >= 0 {
		l.index--
	}
}

func (l *Lexer[T]) LookBack(n int) ([]rune, error) {
	if l.index-n < 0 {
		return nil, io.EOF
	}

	return l.runes[l.index-n : l.index+1], nil
}

// Current returns the current rune at the lexer's index.
func (l *Lexer[T]) Current() rune {
	if l.index >= len(l.runes) {
		return ' '
	}

	return l.runes[l.index]
}

func handleConsumeErr[T comparable](err error, args LexerStateArgs[T]) (LexerStateArgs[T], fsm.State[LexerStateArgs[T]], error) {
	if err != nil {
		if err == io.EOF {
			args.CurrentToken = &shared.Token[T]{
				Type:  args.eofTokenType,
				Value: nil,
			}
		}

		return args, nil, err
	}

	return args, nil, nil
}

func startState[T comparable](ctx context.Context, args LexerStateArgs[T]) (LexerStateArgs[T], fsm.State[LexerStateArgs[T]], error) {
	initialChar, err := args.lexer.Consume()
	args, _, err = handleConsumeErr(err, args)

	if err != nil {
		return args, nil, err
	}

	matchedRule, err := args.lexer.ruleSet.GetMatchingRule(initialChar, args.lexer)
	if err != nil {
		return args, nil, err
	}

	args.currentBuffer.WriteRune(initialChar)
	fn := args.lexer.stateMap[matchedRule]
	return fn(ctx, args)
}

func (l *Lexer[T]) GetNextToken() *shared.Token[T] {
	//fmt.Println("Getting token...")

	args, err := fsm.Run(context.Background(), LexerStateArgs[T]{
		lexer:         l,
		CurrentToken:  nil,
		currentBuffer: &bytes.Buffer{},
		eofTokenType:  l.eofTokenType,
	}, startState)

	if err != nil && err != io.EOF {
		panic(err)
	}

	return args.CurrentToken
}

// GetTokens retrieves all tokens from the lexer's input.
func (l *Lexer[T]) GetTokens() []*shared.Token[T] {
	tokens := make([]*shared.Token[T], 0)

	for {
		token := l.GetNextToken()

		if token == nil {
			continue
		}

		if token.Type == l.eofTokenType {
			break
		}

		tokens = append(tokens, token)
	}

	return tokens
}

// Reset resets the lexer's index and buffer.
func (l *Lexer[T]) Reset() {
	l.index = -1
}
