package lexing

type TokenType int

const (
	InvalidToken TokenType = iota
	EOFToken
	NumberToken
	WhitespaceToken
)
