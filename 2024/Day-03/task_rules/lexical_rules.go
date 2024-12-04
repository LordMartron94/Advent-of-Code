package task_rules

import (
	"bytes"
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/scanning"
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
	DoKeywordToken
	DontKeywordToken
)

type MulKeywordRuleLex struct {
	buffer *bytes.Buffer
}

func (k *MulKeywordRuleLex) IsMatch(scanner scanning.PeekInterface) bool {
	if k.buffer == nil {
		k.buffer = &bytes.Buffer{}
	}

	runesInBuffer, err := scanner.Peek(3)

	if err != nil {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-3] != 'm' {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-2] != 'u' {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-1] != 'l' {
		return false
	}

	for _, r := range runesInBuffer {
		k.buffer.WriteRune(r)
	}

	return true
}

func (k *MulKeywordRuleLex) ExtractToken() (*shared.Token[LexingTokenType], error, int) {
	if k.buffer.Len() < 3 {
		return nil, fmt.Errorf("invalid token: %s", k.buffer.String()), 0
	}

	t := &shared.Token[LexingTokenType]{
		Type:  MulKeywordToken,
		Value: k.buffer.Bytes(),
	}

	k.buffer.Reset()

	return t, nil, 3
}

func (k *MulKeywordRuleLex) Symbol() string {
	return "MulKeywordRuleLex"
}

type DoKeywordRuleLex struct {
	buffer *bytes.Buffer
}

func (d *DoKeywordRuleLex) IsMatch(scanner scanning.PeekInterface) bool {
	if d.buffer == nil {
		d.buffer = &bytes.Buffer{}
	}

	runesInBuffer, err := scanner.Peek(2)

	if err != nil {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-2] != 'd' {
		return false
	}

	if runesInBuffer[len(runesInBuffer)-1] != 'o' {
		return false
	}

	for _, r := range runesInBuffer {
		d.buffer.WriteRune(r)
	}

	return true
}

func (d *DoKeywordRuleLex) ExtractToken() (*shared.Token[LexingTokenType], error, int) {
	if d.buffer.Len() < 2 {
		return nil, fmt.Errorf("invalid token: %s", d.buffer.String()), 0
	}

	t := &shared.Token[LexingTokenType]{
		Type:  DoKeywordToken,
		Value: d.buffer.Bytes(),
	}

	d.buffer.Reset()

	return t, nil, 2
}

func (d *DoKeywordRuleLex) Symbol() string {
	return "DoKeywordRuleLex"
}

type DontKeywordRuleLex struct {
	buffer *bytes.Buffer
}

func (d *DontKeywordRuleLex) IsMatch(scanner scanning.PeekInterface) bool {
	if d.buffer == nil {
		d.buffer = &bytes.Buffer{}
	}

	runesInBuffer, err := scanner.Peek(5)

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

	if runesInBuffer[len(runesInBuffer)-1] != 't' {
		return false
	}

	for _, r := range runesInBuffer {
		d.buffer.WriteRune(r)
	}

	return true
}

func (d *DontKeywordRuleLex) ExtractToken() (*shared.Token[LexingTokenType], error, int) {
	if d.buffer.Len() < 5 {
		return nil, fmt.Errorf("invalid token: %s", d.buffer.String()), 0
	}

	t := &shared.Token[LexingTokenType]{
		Type:  DontKeywordToken,
		Value: d.buffer.Bytes(),
	}

	d.buffer.Reset()

	return t, nil, 5
}

func (d *DontKeywordRuleLex) Symbol() string {
	return "DontKeywordRuleLex"
}

type OpenParenthesisRuleLex struct {
	buffer *bytes.Buffer
}

func (o *OpenParenthesisRuleLex) IsMatch(scanner scanning.PeekInterface) bool {
	if o.buffer == nil {
		o.buffer = &bytes.Buffer{}
	}

	runes, err := scanner.Peek(1)
	r := runes[0]

	if err != nil {
		return false
	}

	if r != '(' {
		return false
	}

	for _, r := range runes {
		o.buffer.WriteRune(r)
	}

	return true
}

func (o *OpenParenthesisRuleLex) ExtractToken() (*shared.Token[LexingTokenType], error, int) {
	if o.buffer.Len() < 1 {
		return nil, fmt.Errorf("invalid token: %s", o.buffer.String()), 0
	}

	t := &shared.Token[LexingTokenType]{
		Type:  OpenParenthesisToken,
		Value: o.buffer.Bytes(),
	}

	o.buffer.Reset()

	return t, nil, 1
}

func (o *OpenParenthesisRuleLex) Symbol() string {
	return "OpenParenthesisRuleLex"
}

type CloseParenthesisRuleLex struct {
	buffer *bytes.Buffer
}

func (c *CloseParenthesisRuleLex) IsMatch(scanner scanning.PeekInterface) bool {
	if c.buffer == nil {
		c.buffer = &bytes.Buffer{}
	}

	runes, err := scanner.Peek(1)
	r := runes[0]

	if err != nil {
		return false
	}

	if r != ')' {
		return false
	}

	for _, r := range runes {
		c.buffer.WriteRune(r)
	}

	return true
}

func (c *CloseParenthesisRuleLex) ExtractToken() (*shared.Token[LexingTokenType], error, int) {
	if c.buffer.Len() < 1 {
		return nil, fmt.Errorf("invalid token: %s", c.buffer.String()), 0
	}

	t := &shared.Token[LexingTokenType]{
		Type:  CloseParenthesisToken,
		Value: c.buffer.Bytes(),
	}

	c.buffer.Reset()

	return t, nil, 1
}

func (c *CloseParenthesisRuleLex) Symbol() string {
	return "CloseParenthesisRuleLex"
}

type CommaRuleLex struct {
	buffer *bytes.Buffer
}

func (c *CommaRuleLex) IsMatch(scanner scanning.PeekInterface) bool {
	if c.buffer == nil {
		c.buffer = &bytes.Buffer{}
	}

	runes, err := scanner.Peek(1)
	r := runes[0]

	if err != nil {
		return false
	}

	if r != ',' {
		return false
	}

	for _, r := range runes {
		c.buffer.WriteRune(r)
	}

	return true
}

func (c *CommaRuleLex) ExtractToken() (*shared.Token[LexingTokenType], error, int) {
	if c.buffer.Len() < 1 {
		return nil, fmt.Errorf("invalid token: %s", c.buffer.String()), 0
	}

	t := &shared.Token[LexingTokenType]{
		Type:  CommaToken,
		Value: c.buffer.Bytes(),
	}

	c.buffer.Reset()

	return t, nil, 1
}

func (c *CommaRuleLex) Symbol() string {
	return "CommaRuleLex"
}

type DigitRuleLex struct {
	buffer *bytes.Buffer
}

func (d *DigitRuleLex) IsMatch(scanner scanning.PeekInterface) bool {
	if d.buffer == nil {
		d.buffer = &bytes.Buffer{}
	}

	runes, err := scanner.Peek(1)
	r := runes[0]

	if err != nil {
		return false
	}

	if r < '0' || r > '9' {
		return false
	}

	for _, r := range runes {
		d.buffer.WriteRune(r)
	}

	return true
}

func (d *DigitRuleLex) ExtractToken() (*shared.Token[LexingTokenType], error, int) {
	if d.buffer.Len() == 0 {
		return nil, fmt.Errorf("invalid token: %s", d.buffer.String()), 0
	}

	t := &shared.Token[LexingTokenType]{
		Type:  NumberToken,
		Value: d.buffer.Bytes(),
	}
	n := d.buffer.Len()

	d.buffer.Reset()

	return t, nil, n
}

func (d *DigitRuleLex) Symbol() string {
	return "DigitRuleLex"
}

type InvalidTokenLex struct {
	buffer *bytes.Buffer
}

func (i *InvalidTokenLex) WriteRune(peeker scanning.PeekInterface) {
	if i.buffer == nil {
		i.buffer = &bytes.Buffer{}
	}

	runes, err := peeker.Peek(1)
	if err != nil {
		return
	}
	r := runes[0]

	i.buffer.WriteRune(r)
}

func (i *InvalidTokenLex) IsMatch(_ scanning.PeekInterface) bool {
	return true
}

func (i *InvalidTokenLex) ExtractToken() (*shared.Token[LexingTokenType], error, int) {
	t := &shared.Token[LexingTokenType]{
		Type:  InvalidToken,
		Value: append([]byte(nil), i.buffer.Bytes()...),
	}
	n := i.buffer.Len()

	i.buffer.Reset()

	return t, nil, n
}

func (i *InvalidTokenLex) Symbol() string {
	return "InvalidTokenLex"
}
