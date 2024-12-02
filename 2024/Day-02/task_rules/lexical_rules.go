package task_rules

import (
	"bytes"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type SpaceRuleLex struct {
}

func (w *SpaceRuleLex) Match(r rune) bool {
	m := r == ' '

	return m
}

func (w *SpaceRuleLex) CreateToken(_ *bytes.Buffer) *shared.Token {
	return &shared.Token{Type: shared.SpaceToken, Value: nil}
}

func (w *SpaceRuleLex) GetName() string {
	return "Space"
}

type NewlineRuleLex struct {
}

func (w *NewlineRuleLex) Match(r rune) bool {
	m := r == '\n'

	return m
}

func (w *NewlineRuleLex) CreateToken(_ *bytes.Buffer) *shared.Token {
	return &shared.Token{Type: shared.NewLineToken, Value: nil}
}

func (w *NewlineRuleLex) GetName() string {
	return "Newline"
}
