package utilities

import (
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/default_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

// FileHandler is a utility struct for Advent of Code file handling.
type FileHandler struct {
	lexer  *lexing.Lexer
	parser *parsing.Parser
}

func NewFileHandler(reader io.Reader, lexingRules []default_rules.LexingRuleInterface, parsingRules []rules.ParsingRuleInterface) *FileHandler {
	lexer := lexing.NewLexer(reader, lexingRules)

	return &FileHandler{
		lexer:  lexer,
		parser: parsing.NewParser(lexer, parsingRules)}
}

func (fh *FileHandler) Lex() []*shared.Token {
	return fh.lexer.GetTokens()
}

func (fh *FileHandler) Parse() (*shared2.ParseTree, error) {
	return fh.parser.Parse()
}

func (fh *FileHandler) ResetLexer() {
	fh.lexer.Reset()
}
