package rules

type LexerInterface interface {
	Peek() (rune, error)
	PeekN(n int) ([]rune, error)
	Consume() (rune, error)
	Pushback()
	LookBack(n int) ([]rune, error)
}
