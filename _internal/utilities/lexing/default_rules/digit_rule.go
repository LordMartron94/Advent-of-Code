package default_rules

import (
	"bytes"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type DigitRule struct {
}

func (d *DigitRule) Match(r rune) bool {
	m := r >= '0' && r <= '9'

	return m
}

func (d *DigitRule) CreateToken(buffer *bytes.Buffer) *shared.Token {
	return &shared.Token{Type: shared.NumberToken, Value: buffer.Bytes()}
}

func (d *DigitRule) GetName() string {
	return "DigitRule"
}
