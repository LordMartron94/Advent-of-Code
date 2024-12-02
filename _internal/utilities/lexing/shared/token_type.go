package shared

type TokenType int

const (
	InvalidToken TokenType = iota
	EOFToken
	NumberToken
	WhitespaceToken
)
