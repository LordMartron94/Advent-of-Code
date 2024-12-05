package common

import (
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type PipelineContext[T comparable] struct {
	Reader                     io.Reader
	ParseTree                  *shared.ParseTree[T]
	Column1Slice, Column2Slice []int
	Distances                  []int
	TotalDistance              int
	TotalIncreases             int
}

func NewPipelineContext[T comparable](openedFile io.Reader) *PipelineContext[T] {
	return &PipelineContext[T]{
		Reader:         openedFile,
		ParseTree:      nil,
		Column1Slice:   make([]int, 0),
		Column2Slice:   make([]int, 0),
		Distances:      make([]int, 0),
		TotalDistance:  0,
		TotalIncreases: 0,
	}
}
