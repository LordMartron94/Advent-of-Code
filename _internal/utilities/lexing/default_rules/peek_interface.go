package default_rules

type LexerInterface interface {
	Peek() (rune, error)
	PeekN(n int) ([]rune, error)
	Consume() (rune, error)
	Pushback()
	LookBack(n int) ([]rune, error)
}

type GetRuleSetInterface[T any] interface {
	GetRuleset() Ruleset[T]
}
