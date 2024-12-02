package default_rules

import (
	"bytes"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type WhitespaceRule struct {
}

func (w *WhitespaceRule) Match(r rune) bool {
	m := r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\f' || r == '\v'

	return m
}

func (w *WhitespaceRule) CreateToken(_ *bytes.Buffer) *shared.Token {
	return &shared.Token{Type: shared.WhitespaceToken, Value: nil}
}

func (w *WhitespaceRule) GetName() string {
	return "Whitespace"
}
