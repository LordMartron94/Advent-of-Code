package default_rules

import (
	"bytes"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type LexingRule interface {
	Match(rune) bool
	CreateToken(buffer *bytes.Buffer) *shared.Token
	GetName() string
}
