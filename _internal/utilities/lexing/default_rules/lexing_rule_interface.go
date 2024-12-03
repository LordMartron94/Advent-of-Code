package default_rules

import (
	"bytes"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type LexingRuleInterface[T any] interface {
	Match(rune, LexerInterface) bool
	CreateToken(buffer *bytes.Buffer) *shared.Token[T]
	GetName() string
}
