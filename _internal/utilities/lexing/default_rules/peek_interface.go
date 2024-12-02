package default_rules

type PeekInterface interface {
	Peek() (rune, error)
	GetRuleset() Ruleset
}
