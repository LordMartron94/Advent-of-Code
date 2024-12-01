package lexing

import (
	"bufio"
	"bytes"
	"io"
)

// Lexer is a struct for lexing input.
type Lexer struct {
	runes []rune
	index int
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

// GetNextToken returns the next token from the input.
func (l *Lexer) GetNextToken() Token {
	return l.StartState()
}

func (l *Lexer) StartState() Token {
	for {
		cRune, err := l.Consume()

		if err == io.EOF {
			return Token{Type: EOFToken, Value: nil}
		}

		if runeIsWhiteSpace(cRune) {
			return l.WhitespaceState()
		}
		if runeIsDigit(cRune) {
			return l.NumberState()
		}
	}
}

func (l *Lexer) WhitespaceState() Token {
	return Token{Type: WhitespaceToken, Value: nil}
}

func (l *Lexer) NumberState() Token {
	buf := bytes.Buffer{}

	buf.WriteRune(l.Current())

	for {
		cRune, err := l.Consume()

		if err == io.EOF {
			return Token{Type: NumberToken, Value: buf.Bytes()}
		}

		if runeIsDigit(cRune) {
			buf.WriteRune(cRune)
			continue
		}
		if runeIsWhiteSpace(cRune) {
			l.Pushback()
			return Token{Type: NumberToken, Value: buf.Bytes()}
		}
	}
}

func runeIsWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\f' || r == '\v'
}

func runeIsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
