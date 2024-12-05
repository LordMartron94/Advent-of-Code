package common

import (
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type PipelineContext[T comparable] struct {
	Reader      io.Reader
	ParseTree   *shared.ParseTree[T]
	Result      int
	Manuals     [][]int
	Updates     [][]int
	FixedResult int
}

func NewPipelineContext[T comparable](openedFile io.Reader) *PipelineContext[T] {
	return &PipelineContext[T]{
		Reader:      openedFile,
		ParseTree:   nil,
		Result:      0,
		Manuals:     make([][]int, 0),
		Updates:     make([][]int, 0),
		FixedResult: 0,
	}
}
