package task_rules

import (
	"bytes"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/default_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type LexingTokenType int

const (
	InvalidToken LexingTokenType = iota
	MulKeywordToken
	OpenParenthesisToken
	CloseParenthesisToken
	CommaToken
	NumberToken
	EOFToken
	DoKeywordToken
	DontKeywordToken
)

type MulKeywordRuleLex struct {
}

func (k *MulKeywordRuleLex) Match(r rune, lexer default_rules.LexerInterface) bool {
	runesInBuffer, err := lexer.LookBack(3)

	if err != nil {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-3] != 'm' {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-2] != 'u' {
		return false
	}

	if r != 'l' {
		return false
	}

	return true
}

func (k *MulKeywordRuleLex) CreateToken(buffer *bytes.Buffer) *shared.Token[LexingTokenType] {
	return &shared.Token[LexingTokenType]{
		Type:  MulKeywordToken,
		Value: []byte("mul"),
	}
}

func (k *MulKeywordRuleLex) GetName() string {
	return "MulKeywordRuleLex"
}

type DoKeywordRuleLex struct{}

func (d *DoKeywordRuleLex) Match(r rune, lexer default_rules.LexerInterface) bool {
	runesInBuffer, err := lexer.LookBack(2)

	if err != nil {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-2] != 'd' {
		return false
	}

	if r != 'o' {
		return false
	}

	return true
}

func (d *DoKeywordRuleLex) CreateToken(buffer *bytes.Buffer) *shared.Token[LexingTokenType] {
	return &shared.Token[LexingTokenType]{
		Type:  DoKeywordToken,
		Value: []byte("do"),
	}
}

func (d *DoKeywordRuleLex) GetName() string {
	return "DoKeywordRuleLex"
}

type DontKeywordRuleLex struct{}

func (d *DontKeywordRuleLex) Match(r rune, lexer default_rules.LexerInterface) bool {
	runesInBuffer, err := lexer.LookBack(5)

	if err != nil {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-5] != 'd' {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-4] != 'o' {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-3] != 'n' {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-2] != '\'' {
		return false
	}

	if r != 't' {
		return false
	}

	return true
}

func (d *DontKeywordRuleLex) CreateToken(buffer *bytes.Buffer) *shared.Token[LexingTokenType] {
	return &shared.Token[LexingTokenType]{
		Type:  DontKeywordToken,
		Value: []byte("don't"),
	}
}

func (d *DontKeywordRuleLex) GetName() string {
	return "DontKeywordRuleLex"
}

type OpenParenthesisRuleLex struct{}

func (o *OpenParenthesisRuleLex) Match(r rune, lexer default_rules.LexerInterface) bool {
	return r == '('
}

func (o *OpenParenthesisRuleLex) CreateToken(buffer *bytes.Buffer) *shared.Token[LexingTokenType] {
	return &shared.Token[LexingTokenType]{
		Type:  OpenParenthesisToken,
		Value: buffer.Bytes(),
	}
}

func (o *OpenParenthesisRuleLex) GetName() string {
	return "OpenParenthesisRuleLex"
}

type CloseParenthesisRuleLex struct{}

func (c *CloseParenthesisRuleLex) Match(r rune, lexer default_rules.LexerInterface) bool {
	return r == ')'
}

func (c *CloseParenthesisRuleLex) CreateToken(buffer *bytes.Buffer) *shared.Token[LexingTokenType] {
	return &shared.Token[LexingTokenType]{
		Type:  CloseParenthesisToken,
		Value: buffer.Bytes(),
	}
}

func (c *CloseParenthesisRuleLex) GetName() string {
	return "CloseParenthesisRuleLex"
}

type CommaRuleLex struct{}

func (c *CommaRuleLex) Match(r rune, lexer default_rules.LexerInterface) bool {
	return r == ','
}

func (c *CommaRuleLex) CreateToken(buffer *bytes.Buffer) *shared.Token[LexingTokenType] {
	return &shared.Token[LexingTokenType]{
		Type:  CommaToken,
		Value: buffer.Bytes(),
	}
}

func (c *CommaRuleLex) GetName() string {
	return "CommaRuleLex"
}

type DigitRuleLex struct{}

func (d *DigitRuleLex) Match(r rune, lexer default_rules.LexerInterface) bool {
	return '0' <= r && r <= '9'
}

func (d *DigitRuleLex) CreateToken(buffer *bytes.Buffer) *shared.Token[LexingTokenType] {
	return &shared.Token[LexingTokenType]{
		Type:  NumberToken,
		Value: buffer.Bytes(),
	}
}

func (d *DigitRuleLex) GetName() string {
	return "DigitRuleLex"
}

type InvalidTokenLex struct{}

func (i *InvalidTokenLex) Match(_ rune, lexer default_rules.LexerInterface) bool {
	return true
}

func (i *InvalidTokenLex) CreateToken(buffer *bytes.Buffer) *shared.Token[LexingTokenType] {
	return &shared.Token[LexingTokenType]{
		Type:  InvalidToken,
		Value: buffer.Bytes(),
	}
}

func (i *InvalidTokenLex) GetName() string {
	return "InvalidTokenLex"
}
