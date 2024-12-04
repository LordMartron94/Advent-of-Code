package scanning

type ScannerInterface interface {
	Peek(n int) ([]rune, error)
	Consume(n int) ([]rune, error)
	Pushback(n int) error
	Reset()
}

type PeekInterface interface {
	Peek(n int) ([]rune, error)
}
