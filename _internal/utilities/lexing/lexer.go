package lexing

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/fsm"
)

// Lexer is a struct for lexing input.
type Lexer struct {
	runes []rune
	index int
}

type LexerStateArgs struct {
	lexer *Lexer

	CurrentToken  *Token
	currentBuffer *bytes.Buffer
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

	return &Lexer{
		runes: runes,
		index: -1,
	}
}

// Peek returns the next rune without advancing the lexer's index.
func (l *Lexer) Peek() (rune, error) {
	if l.index+1 > len(l.runes) {
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

// GetNextToken retrieves the token from the lexer's input.
func (l *Lexer) GetNextToken() *Token {
	args, err := fsm.Run(context.Background(), LexerStateArgs{
		lexer:         l,
		CurrentToken:  nil,
		currentBuffer: &bytes.Buffer{},
	},
		StartState)

	if err != nil {
		panic(err)
	}

	return args.CurrentToken
}

// GetTokens retrieves all tokens from the lexer's input.
func (l *Lexer) GetTokens() []*Token {
	tokens := make([]*Token, 0)

	for {
		token := l.GetNextToken()

		if token.Type == EOFToken {
			break
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func StartState(_ context.Context, args LexerStateArgs) (LexerStateArgs, fsm.State[LexerStateArgs], error) {
	for {
		cRune, err := args.lexer.Consume()

		if err == io.EOF {
			args.CurrentToken = &Token{Type: EOFToken, Value: nil}
			return args, nil, nil
		}

		if runeIsWhiteSpace(cRune) {
			return args, WhitespaceState, nil
		}
		if runeIsDigit(cRune) {
			return args, NumberState, nil
		}

		return args, nil, fmt.Errorf("unexpected character: %c", cRune)
	}
}

func WhitespaceState(_ context.Context, args LexerStateArgs) (LexerStateArgs, fsm.State[LexerStateArgs], error) {
	args.CurrentToken = &Token{Type: WhitespaceToken, Value: nil}
	return args, nil, nil
}

func NumberState(_ context.Context, args LexerStateArgs) (LexerStateArgs, fsm.State[LexerStateArgs], error) {
	args.currentBuffer.WriteRune(args.lexer.Current())

	for {
		cRune, err := args.lexer.Consume()

		if err == io.EOF {
			args.CurrentToken = &Token{Type: NumberToken, Value: args.currentBuffer.Bytes()}
			return args, nil, nil
		}

		if runeIsDigit(cRune) {
			args.currentBuffer.WriteRune(cRune)
			continue
		} else {
			args.lexer.Pushback()
			args.CurrentToken = &Token{Type: NumberToken, Value: args.currentBuffer.Bytes()}
			return args, nil, nil
		}
	}
}

func runeIsWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\f' || r == '\v'
}

func runeIsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
