package utilities

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/default_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

// FileHandler is a utility struct for Advent of Code file handling.
type FileHandler[T comparable] struct {
	lexer  *lexing.Lexer[T]
	parser *parsing.Parser[T]
}

func ChangeWorkingDirectoryToTodayTask() {
	today := time.Now()
	year := today.Year()
	day := fmt.Sprintf("%02d", today.Day())

	err := os.Chdir(fmt.Sprintf("./%d/Day-%s", year, day))
	if err != nil {
		fmt.Printf("Error changing working directory to today's task: %v\n", err)
		return
	}
}

func ChangeWorkingDirectoryToSpecificTask(year int, day int) {
	sDay := fmt.Sprintf("%02d", day)

	err := os.Chdir(fmt.Sprintf("./%d/Day-%s", year, sDay))
	if err != nil {
		fmt.Printf("Error changing working directory to today's task: %v\n", err)
		return
	}
}

func NewFileHandler[T comparable](reader io.Reader, lexingRules []default_rules.LexingRuleInterface[T], parsingRules []rules.ParsingRuleInterface[T], eofTokenType T) *FileHandler[T] {
	lexer := lexing.NewLexer[T](reader, lexingRules, eofTokenType)

	return &FileHandler[T]{
		lexer:  lexer,
		parser: parsing.NewParser[T](lexer, parsingRules)}
}

func (fh *FileHandler[T]) Lex() []*shared.Token[T] {
	return fh.lexer.GetTokens()
}

func (fh *FileHandler[T]) Parse() (*shared2.ParseTree[T], error) {
	return fh.parser.Parse()
}

func (fh *FileHandler[T]) ResetLexer() {
	fh.lexer.Reset()
}
