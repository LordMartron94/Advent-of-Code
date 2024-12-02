package common

import (
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type PipelineContext struct {
	Reader                     io.Reader
	ParseTree                  *shared.ParseTree
	Column1Slice, Column2Slice []int
	Distances                  []int
	TotalDistance              int
	TotalIncreases             int
}

func NewPipelineContext(openedFile io.Reader) *PipelineContext {
	return &PipelineContext{
		Reader:         openedFile,
		ParseTree:      nil,
		Column1Slice:   make([]int, 0),
		Column2Slice:   make([]int, 0),
		Distances:      make([]int, 0),
		TotalDistance:  0,
		TotalIncreases: 0,
	}
}
