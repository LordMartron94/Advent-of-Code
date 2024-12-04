package lexing

import "github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"

type LexerInterface[T comparable] interface {
	// GetToken returns the next token from the input stream.
	GetToken() *shared.Token[T]

	// GetTokens returns all tokens from the input stream.
	GetTokens() ([]*shared.Token[T], error)

	// Reset resets the lexer to the beginning of the input stream.
	Reset()
}
