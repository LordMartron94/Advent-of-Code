package shared

import "fmt"

// Token represents a lexical token
type Token[T any] struct {
	Type  T
	Value []byte
}

func (t *Token[T]) String() string {
	return fmt.Sprintf("[(%v) - '%s']", t.Type, t.Value)
}

func TokensToStrings[T comparable](tokens []Token[T]) []string {
	stringsToReturn := make([]string, 0)
	for _, token := range tokens {
		stringsToReturn = append(stringsToReturn, fmt.Sprintf("%s ", token.String()))
	}
	return stringsToReturn
}
