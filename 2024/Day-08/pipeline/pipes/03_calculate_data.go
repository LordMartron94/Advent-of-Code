package pipes

import (
	"slices"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-08/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-08/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type CalculateDataPipe struct{}

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	rows := input.Rows

	basicComparer := func(a, b shared.Token[task_rules.LexingTokenType]) bool {
		return a.Equals(b)
	}

	matrixHelper := matrix.NewMatrixHelper(rows, basicComparer)

	customComparer := func(a, b shared.Token[task_rules.LexingTokenType]) bool {
		return a.Type == b.Type
	}

	antennaTypes, antennaPositions := matrixHelper.GetCoordinatesOfTypesFiltered([]shared.Token[task_rules.LexingTokenType]{{Type: task_rules.AlphanumericToken}}, &customComparer, basicComparer)
	antiNodesPlaced := make([]matrix.Position, 0)
	antiNodesPlacedPart2 := make([]matrix.Position, 0)

	for i := range antennaTypes {
		antennasForTypePositions := antennaPositions[i]
		if len(antennasForTypePositions) < 2 {
			continue
		}

		distances := matrixHelper.AggregateUniqueDistancesBetweenPositions(antennasForTypePositions)

		for _, distance := range distances {
			firstAntiNodePos, secondAntiNodePos := matrixHelper.GetExtendedLinePositions(distance)
			antiNodesPart2 := matrixHelper.GetLinePositions(distance)

			if firstAntiNodePos != nil {
				if !slices.Contains(antiNodesPlaced, *firstAntiNodePos) {
					antiNodesPlaced = append(antiNodesPlaced, *firstAntiNodePos)
				}
			}

			if secondAntiNodePos != nil {
				if !slices.Contains(antiNodesPlaced, *secondAntiNodePos) {
					antiNodesPlaced = append(antiNodesPlaced, *secondAntiNodePos)
				}
			}

			for _, pos := range antiNodesPart2 {
				if !slices.Contains(antiNodesPlacedPart2, pos) {
					antiNodesPlacedPart2 = append(antiNodesPlacedPart2, pos)
				}
			}
		}
	}

	input.Result = len(antiNodesPlaced)
	input.ResultPart2 = len(antiNodesPlacedPart2)

	return input
}
