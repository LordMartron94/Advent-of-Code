package pipes

import (
	"fmt"
	"sort"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-05/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-05/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
)

type CalculateDataPipe struct {
}

func sliceContainsAllValues[T comparable](slice []T, values []T) bool {
	for _, value2 := range values {
		found := false
		for _, value1 := range slice {
			if value1 == value2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

type Update struct {
	UpdateContents    []int
	AssociatedManuals [][]int
	FixedManual       []int
	Correct           bool
}

func (c *CalculateDataPipe) GetManualUpdateCombinations(manuals [][]int, updates [][]int) []Update {
	combinations := make([]Update, 0)

	for _, update := range updates {
		combination := Update{
			UpdateContents:    update,
			AssociatedManuals: make([][]int, 0),
			FixedManual:       make([]int, 0),
		}

		for _, manual := range manuals {
			if sliceContainsAllValues(update, manual) {
				combination.AssociatedManuals = append(combination.AssociatedManuals, manual)
			}
		}

		combinations = append(combinations, combination)
	}

	return combinations
}

func (c *CalculateDataPipe) MergeManuals(manualPairs [][]int) []int {
	graphHelper := transforming.NewGraphHelper()

	sorted := graphHelper.TopologicalSortPairs(manualPairs)

	//fmt.Println("Sorted manual pairs:", sorted)
	//os.Exit(0)

	return sorted
}

func (c *CalculateDataPipe) FixManuals(combinations []Update) {
	for i := range combinations {
		combinations[i].FixedManual = c.MergeManuals(combinations[i].AssociatedManuals)
	}
}

func (c *CalculateDataPipe) AreManualsCorrect(manual []int, update []int) bool {
	toCheckValues := make([]int, 0)

	for i := 0; i < len(update); i++ {
		updateValue := update[i]

		for _, manualValue := range manual {
			if updateValue == manualValue {
				toCheckValues = append(toCheckValues, updateValue)
				break
			}
		}
	}

	matching := true

	if len(toCheckValues) != len(manual) {
		fmt.Println("Mismatch in manuals for update:", update)
		return false
	}

	for i := range toCheckValues {
		if toCheckValues[i] != manual[i] {
			matching = false
			break
		}
	}

	return matching
}

func (c *CalculateDataPipe) GetCorrectUpdates(updates []Update) {
	for i := range updates {
		updates[i].Correct = c.AreManualsCorrect(updates[i].FixedManual, updates[i].UpdateContents)
	}
}

func sortUpdate(update []int, rules []int) []int {
	// Create a map to store the order of pages based on the rules
	orderMap := make(map[int]int)
	for i, page := range rules {
		orderMap[page] = i
	}

	// Sort the update based on the order defined in the rules
	sort.Slice(update, func(j, k int) bool {
		return orderMap[update[j]] < orderMap[update[k]]
	})

	return update
}

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	combinations := c.GetManualUpdateCombinations(input.Manuals, input.Updates)
	c.FixManuals(combinations)
	c.GetCorrectUpdates(combinations)

	sumMiddlePage := 0
	correctedSumMiddlePage := 0

	for _, combination := range combinations {
		if combination.Correct {
			if len(combination.UpdateContents)%2 != 0 {
				sumMiddlePage += combination.UpdateContents[len(combination.UpdateContents)/2]
			} else {
				fmt.Println("Even encountered, skoippp")
				continue
			}
		} else {
			updated := sortUpdate(combination.UpdateContents, combination.FixedManual)
			if len(updated)%2 != 0 {
				correctedSumMiddlePage += updated[len(updated)/2]
			} else {
				fmt.Println("Even encountered, skoippp")
				continue
			}
		}
	}

	input.Result = sumMiddlePage
	input.FixedResult = correctedSumMiddlePage

	return input
}
