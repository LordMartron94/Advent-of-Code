package common

import (
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type Equation struct {
	TestNumber int
	Elements   []int
}

type PipelineContext[T comparable] struct {
	Reader     io.Reader
	ParseTree  *shared.ParseTree[T]
	Result     int
	Equations  []Equation
	ResultPipe int
}

func NewPipelineContext[T comparable](openedFile io.Reader) *PipelineContext[T] {
	return &PipelineContext[T]{
		Reader:     openedFile,
		ParseTree:  nil,
		Result:     0,
		Equations:  make([]Equation, 0),
		ResultPipe: 0,
	}
}
