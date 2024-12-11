package pipes

import (
	"fmt"
	"math"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-11/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-11/task_rules"
)

type CalculateDataPipe struct {
	stoneCache map[int][2]int
}

func digitCount(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(float64(n)))) + 1
}

func splitNumber(n int, numDigits int) (int, int) {
	divisor := int(math.Pow10(numDigits / 2))
	return n / divisor, n % divisor
}

//if stone == 0 {
//return []int{1}
//}
//
//numDigits := digitCount(stone)
//
//if numDigits&1 == 0 {
//leftHalf, rightHalf := splitNumber(stone, numDigits)
//return []int{leftHalf, rightHalf}
//}
//
//return []int{stone * 2024}

// The tricky part is now looking at the problem more closely and trying to find the hidden structure in the rules that lets you make the problem tractable.
// (In a programming contest, this always exists, but not necessarily so in the real world.)
// Given this puzzle's rules, you can figure out that you only need to keep track of a certain stone number for each label and figure out the algorithm for creating new counts corresponding to a round of expansion.
// ^ Dale

func getResultForStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	numDigits := digitCount(stone)

	if numDigits&1 == 0 {
		leftHalf, rightHalf := splitNumber(stone, numDigits)
		return []int{leftHalf, rightHalf}
	}

	return []int{stone * 2024}
}

func (c *CalculateDataPipe) addValueToMap(m map[int]int, values []int, count int) {
	for _, value := range values {
		if _, exists := m[value]; !exists {
			m[value] = count
		} else {
			m[value] += count
		}
	}
}

func (c *CalculateDataPipe) transformStones(stones map[int]int) map[int]int {
	newMap := make(map[int]int)
	for value, count := range stones {
		newValue := getResultForStone(value)
		c.addValueToMap(newMap, newValue, count)
	}

	return newMap
}

func (c *CalculateDataPipe) getStoneMap(stones []int) map[int]int {
	stoneMap := make(map[int]int)
	for _, stone := range stones {
		if _, exists := stoneMap[stone]; !exists {
			stoneMap[stone] = 1
		} else {
			stoneMap[stone]++
		}
	}

	return stoneMap
}

func (c *CalculateDataPipe) GetStoneCount(input common.PipelineContext[task_rules.LexingTokenType], numIterations int) int {
	initialStoneMap := c.getStoneMap(input.Stones)
	transformedStoneMap := initialStoneMap

	for i := 0; i < numIterations; i++ {
		fmt.Println("Iteration: ", i+1)
		transformedStoneMap = c.transformStones(transformedStoneMap)
	}

	totalCount := 0

	for _, count := range transformedStoneMap {
		totalCount += count
	}

	return totalCount
}

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	input.Result = c.GetStoneCount(input, 25)
	fmt.Println("-----------------------")
	input.Result2 = c.GetStoneCount(input, 75)
	return input
}
