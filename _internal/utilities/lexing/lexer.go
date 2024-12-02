package lexing

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/fsm"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/default_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type LexerStateArgs struct {
	lexer *Lexer

	CurrentToken  *shared.Token
	currentBuffer *bytes.Buffer
}

// Lexer is a struct for lexing input.
type Lexer struct {
	runes []rune
	index int

	ruleSet  *default_rules.Ruleset
	stateMap map[default_rules.LexingRule]fsm.State[LexerStateArgs]
}

// NewLexer creates a new Lexer with the given reader.
func NewLexer(reader io.Reader) *Lexer {
	bReader := bufio.NewReader(reader)

	runes := make([]rune, 0)

	for {
		scannedRune, _, err := bReader.ReadRune()

		if err == io.EOF {
			break
		}

		runes = append(runes, scannedRune)
	}

	lexer := &Lexer{
		runes: runes,
		index: -1,
	}

	rules := make([]default_rules.LexingRule, 0)
	rules = append(rules, &default_rules.WhitespaceRule{})
	rules = append(rules, &default_rules.DigitRule{})

	lexer.ruleSet = default_rules.NewRuleset(rules)

	sM, err := lexer.generateFSM()

	if err != nil {
		panic(err)
	}

	lexer.stateMap = sM

	return lexer
}

func (l *Lexer) GetRuleset() default_rules.Ruleset {
	return *l.ruleSet
}

func (l *Lexer) generateFSM() (map[default_rules.LexingRule]fsm.State[LexerStateArgs], error) {
	stateMap := make(map[default_rules.LexingRule]fsm.State[LexerStateArgs])

	for _, rule := range l.ruleSet.Rules {
		stateMap[rule] = func(ctx context.Context, args LexerStateArgs) (LexerStateArgs, fsm.State[LexerStateArgs], error) {
			peekRune, err := args.lexer.Peek()
			if err != nil {
				return args, nil, err
			}

			fmt.Println("Current rune: ", string(peekRune))

			return args, nil, nil
		}
	}

	return stateMap, nil
}

// Peek returns the next rune without advancing the lexer's index.
func (l *Lexer) Peek() (rune, error) {
	if l.index+1 >= len(l.runes) {
		return ' ', io.EOF
	}

	return l.runes[l.index+1], nil
}

// Consume returns the next rune and advances the lexer's index.
func (l *Lexer) Consume() (rune, error) {
	if l.index+1 >= len(l.runes) {
		return ' ', io.EOF
	}

	l.index++

	return l.runes[l.index], nil
}

// Pushback reverts the lexer's index to the previous position.
func (l *Lexer) Pushback() {
	if l.index >= 0 {
		l.index--
	}
}

// Current returns the current rune at the lexer's index.
func (l *Lexer) Current() rune {
	if l.index >= len(l.runes) {
		return ' '
	}

	return l.runes[l.index]
}

func handleConsumeErr(err error, args LexerStateArgs) (LexerStateArgs, fsm.State[LexerStateArgs], error) {
	if err != nil {
		if err == io.EOF {
			args.CurrentToken = &shared.Token{
				Type:  shared.EOFToken,
				Value: nil,
			}
		}

		return args, nil, err
	}

	return args, nil, nil
}

func startState(_ context.Context, args LexerStateArgs) (LexerStateArgs, fsm.State[LexerStateArgs], error) {
	initialChar, err := args.lexer.Consume()
	args, _, err = handleConsumeErr(err, args)

	if err != nil {
		return args, nil, err
	}

	matchedRule, err := args.lexer.ruleSet.GetMatchingRule(initialChar)
	if err != nil {
		return args, nil, err
	}

	args.currentBuffer.WriteRune(initialChar)

	for {
		cRune, err := args.lexer.Consume()
		args, _, err = handleConsumeErr(err, args)

		if err != nil {
			if err == io.EOF {
				t := matchedRule.CreateToken(args.currentBuffer)
				args.CurrentToken = t
				return args, nil, err
			}

			return args, nil, err
		}

		matchedRule2, err := args.lexer.ruleSet.GetMatchingRule(cRune)
		if err != nil {
			return args, nil, err
		}

		if matchedRule2 == matchedRule {
			args.currentBuffer.WriteRune(cRune)
			continue
		} else {
			args.lexer.Pushback()
			t := matchedRule.CreateToken(args.currentBuffer)
			args.CurrentToken = t
			return args, nil, nil
		}
	}
}

func (l *Lexer) GetNextToken() *shared.Token {
	args, err := fsm.Run(context.Background(), LexerStateArgs{
		lexer:         l,
		CurrentToken:  nil,
		currentBuffer: &bytes.Buffer{},
	}, startState)

	if err != nil && err != io.EOF {
		panic(err)
	}

	return args.CurrentToken
}

// GetTokens retrieves all tokens from the lexer's input.
func (l *Lexer) GetTokens() []*shared.Token {
	tokens := make([]*shared.Token, 0)

	for {
		token := l.GetNextToken()

		if token == nil {
			continue
		}

		if token.Type == shared.EOFToken {
			break
		}

		tokens = append(tokens, token)
	}

	return tokens
}
