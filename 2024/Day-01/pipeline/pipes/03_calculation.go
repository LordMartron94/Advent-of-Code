package pipes

import (
	"github.com/LordMartron94/Advent-of-Code/2024/Day-01/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-01/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/common_calculations"
)

type CalculationPipe struct {
}

func mapContainsKey(m map[int]int, key int) bool {
	_, exists := m[key]
	return exists
}

func getNumAppearancesInSlice(slice []int, key int) int {
	count := 0

	for _, num := range slice {
		if num == key {
			count++
		}
	}

	return count
}

func constructAppearanceMap(xSlice, ySlice []int) map[int]int {
	appearances := make(map[int]int)

	for _, num := range xSlice {
		if !mapContainsKey(appearances, num) {
			appearances[num] = getNumAppearancesInSlice(ySlice, num)
		} else {
			continue
		}
	}

	return appearances
}

func (c *CalculationPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	totalDistance := 0
	common_calculations.SumInts(&input.Distances, &totalDistance)

	appearanceMap := constructAppearanceMap(input.Column1Slice, input.Column2Slice)
	increases := make([]int, 0)

	common_calculations.MapAndTransformSlice(
		&input.Column1Slice,
		func(num int, appearancesMap map[int]int) int {
			return num * appearancesMap[num]
		},
		appearanceMap,
		&increases,
	)

	totalIncreases := common_calculations.SumIntsAndReturn(increases)

	input.TotalDistance = totalDistance
	input.TotalIncreases = totalIncreases

	return input
}
