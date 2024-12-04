package rules

import (
	"fmt"
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/scanning"
)

type Ruleset[T any] struct {
	Rules []LexingRuleInterface[T]
}

func NewRuleset[T any](rules []LexingRuleInterface[T]) *Ruleset[T] {
	return &Ruleset[T]{Rules: rules}
}

// GetMatchingRule returns the first matching rule for the given input stream.
// If no matching rule is found, it returns an error.
// If the input stream is exhausted before a matching rule is found, it returns io.EOF.
func (rs *Ruleset[T]) GetMatchingRule(scanner scanning.PeekInterface) (LexingRuleInterface[T], error) {
	_, err := scanner.Peek(1)

	if err != nil {
		if err == io.EOF {
			return nil, io.EOF
		}

		return nil, fmt.Errorf("error peeking at scanner: %w", err)
	}

	for _, rule := range rs.Rules {
		matched := rule.IsMatch(scanner)

		if matched {
			//fmt.Println(fmt.Sprintf("Matched rule (ruleSet Matcher): %s for first character '%s'", rule.Symbol(), string(input)))
			return rule, nil
		}
	}

	return nil, fmt.Errorf("no matching rule found\n")
}
