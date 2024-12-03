package shared

// Token represents a lexical token
type Token[T any] struct {
	Type  T
	Value []byte
}
