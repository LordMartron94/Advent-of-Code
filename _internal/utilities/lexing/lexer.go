package lexing

import (
	"fmt"
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/scanning"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

// Lexer is a generic lexer for a given input stream.
type Lexer[T comparable] struct {
	scanner scanning.ScannerInterface
	ruleSet *rules.Ruleset[T]
}

// NewLexer creates a new lexer for the given input stream.
func NewLexer[T comparable](reader io.Reader, lexingRules []rules.LexingRuleInterface[T]) *Lexer[T] {
	scanner := scanning.NewScanner(reader)
	ruleset := rules.NewRuleset[T](lexingRules)

	return &Lexer[T]{
		scanner: scanner,
		ruleSet: ruleset,
	}
}

// GetToken returns the next token from the input stream.
func (l *Lexer[T]) GetToken() *shared.Token[T] {
	matchingRule, err := l.getMatchingRule()
	if err != nil {
		if err == io.EOF {
			return nil
		}

		panic(err)
	}

	if matchingRule == nil {
		panic("no matching rule found")
	}

	consumed, err := l.consumeMatchingRunes(matchingRule)

	if err != nil {
		panic(err)
	}

	return l.extractToken(consumed, matchingRule)
}

// getMatchingRule retrieves the matching rule for the current scanner state.
func (l *Lexer[T]) getMatchingRule() (rules.LexingRuleInterface[T], error) {
	matchingRule, err := l.ruleSet.GetMatchingRule(l.scanner)
	if err != nil {
		if err == io.EOF {
			return nil, io.EOF
		}
		return nil, err
	}
	return matchingRule, nil
}

// consumeMatchingRunes consumes runes while they match the given rule.
func (l *Lexer[T]) consumeMatchingRunes(rule rules.LexingRuleInterface[T]) (int, error) {
	rule.WriteRune(l.scanner)
	consumed := 0
	for {
		if _, err := l.scanner.Consume(1); err != nil {
			if err == io.EOF {
				break
			}
			return consumed, err
		}
		consumed++

		currentRule, err := l.ruleSet.GetMatchingRule(l.scanner)
		if err != nil {
			if err == io.EOF {
				break
			}
			return consumed, err
		}

		if rule != currentRule {
			if err := l.scanner.Pushback(1); err != nil {
				return consumed, err
			}
			consumed--
			break
		} else {
			rule.WriteRune(l.scanner)
		}
	}

	_, _ = l.scanner.Consume(consumed)
	return consumed, nil
}

// extractToken extracts the token from the matched rule.
func (l *Lexer[T]) extractToken(consumedCount int, rule rules.LexingRuleInterface[T]) *shared.Token[T] {
	t, err, consumedN := rule.ExtractToken()
	if err != nil {
		panic(err)
	}
	
	if consumedN != consumedCount {
		// This should theoretically never happen, but it's a good safety net.
		panic(fmt.Errorf("Warning: consumed %d characters but rule matched %d characters\n", consumedCount, consumedN))
	}

	return t
}

// GetTokens returns all tokens from the input stream.
func (l *Lexer[T]) GetTokens() ([]*shared.Token[T], error) {
	tokens := make([]*shared.Token[T], 0)

	for {
		token := l.GetToken()

		if token == nil {
			break
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

// Reset resets the lexer's scanner to its initial state.
func (l *Lexer[T]) Reset() {
	l.scanner.Reset()
}
