package utilities

import (
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing"
)

// FileHandler is a utility struct for Advent of Code file handling.
type FileHandler struct {
	lexer *lexing.Lexer
}

func NewFileHandler(reader io.Reader) *FileHandler {
	return &FileHandler{lexer: lexing.NewLexer(reader)}
}

func (fh *FileHandler) Lex() []*lexing.Token {
	return fh.lexer.GetTokens()
}
