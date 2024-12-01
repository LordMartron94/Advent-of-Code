package lexing

// Token represents a lexical token
type Token struct {
	Type  TokenType
	Value []byte
}
