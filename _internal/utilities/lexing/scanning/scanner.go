package scanning

import (
	"fmt"
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	runes        []rune
	currentIndex int
}

// NewScanner creates a new Scanner with the given input stream.
func NewScanner(reader io.Reader) *Scanner {
	runes, err := shared.ReaderToRunes(reader)

	if err != nil {
		panic(err)
	}

	return &Scanner{
		runes:        runes,
		currentIndex: 0,
	}
}

// Peek returns the next n runes without advancing the scanner's index.
func (s *Scanner) Peek(n int) ([]rune, error) {
	if s.currentIndex+n >= len(s.runes) {
		return nil, io.EOF
	}

	return s.runes[s.currentIndex+1 : s.currentIndex+n+1], nil
}

// Consume returns the next n runes and advances the scanner's index.
func (s *Scanner) Consume(n int) ([]rune, error) {
	runes, err := s.Peek(n)

	if err != nil {
		return nil, err
	}

	s.currentIndex += n

	return runes, nil
}

// Pushback returns the scanner's index by n.
func (s *Scanner) Pushback(n int) error {
	if s.currentIndex-n < 0 {
		return fmt.Errorf("lexer index cannot be pushed back by %d", n)
	}

	s.currentIndex -= n

	return nil
}

// Reset resets the scanner's index to the beginning of the input stream.
func (s *Scanner) Reset() {
	s.currentIndex = 0
}

// Current returns the current rune.
func (s *Scanner) Current() rune {
	if s.currentIndex >= len(s.runes) {
		panic("Lexer index out of range") // Obviously, this should never happen.
	}

	return s.runes[s.currentIndex]
}
