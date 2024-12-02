package common

import (
	"io"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type PipelineContext struct {
	Reader             io.Reader
	ParseTree          *shared.ParseTree
	Reports            [][]int
	SafeReports        int
	SafeReportsRevised int
}

func NewPipelineContext(openedFile io.Reader) *PipelineContext {
	return &PipelineContext{
		Reader:             openedFile,
		ParseTree:          nil,
		Reports:            make([][]int, 0),
		SafeReports:        0,
		SafeReportsRevised: 0,
	}
}
