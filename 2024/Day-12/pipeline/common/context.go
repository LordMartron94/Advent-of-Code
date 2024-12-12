package common

import (
	"io"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-12/task_rules"
	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type PipelineContext[T comparable] struct {
	Reader    io.Reader
	ParseTree *shared.ParseTree[T]
	Result    int
	Result2   int
	Rows      [][]shared3.Token[task_rules.LexingTokenType]
}

func NewPipelineContext[T comparable](openedFile io.Reader) *PipelineContext[T] {
	return &PipelineContext[T]{
		Reader:    openedFile,
		ParseTree: nil,
		Result:    0,
		Result2:   0,
		Rows:      make([][]shared3.Token[task_rules.LexingTokenType], 0),
	}
}
