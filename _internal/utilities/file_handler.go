package utilities

import (
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/default_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

// FileHandler is a utility struct for Advent of Code file handling.
type FileHandler struct {
	lexer *lexing.Lexer
}

func NewFileHandler(reader io.Reader, lexingRules []default_rules.LexingRule) *FileHandler {
	return &FileHandler{lexer: lexing.NewLexer(reader, lexingRules)}
}

func (fh *FileHandler) Lex() []*shared.Token {
	return fh.lexer.GetTokens()
}
