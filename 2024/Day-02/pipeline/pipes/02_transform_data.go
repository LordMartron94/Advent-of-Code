package pipes

import (
	"fmt"
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-02/pipeline/common"
	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

func allItemsIncreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] >= report[i+1] {
			return false
		}
	}

	return true
}

func allItemsDecreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] <= report[i+1] {
			return false
		}
	}

	return true
}

func getAdjacentDistances(report []int) []int { // WOrks
	distances := make([]int, len(report)-1)

	for i := 0; i < len(report)-1; i++ {
		d := report[i+1] - report[i]

		if d < 0 {
			d = -d
		}

		distances[i] = d
	}

	return distances
}

func numWithinRange(target, min, max int) bool { // WOrks
	return target >= min && target <= max
}

func isReportSafe(report []int) bool {
	increasingOrDecreasing := allItemsIncreasing(report) || allItemsDecreasing(report)

	if !increasingOrDecreasing {
		return false
	}

	adjacentDistances := getAdjacentDistances(report)

	safeReport := true

	for _, adjacentDistance := range adjacentDistances {
		if !numWithinRange(adjacentDistance, 1, 3) {
			safeReport = false
			break
		}
	}

	return increasingOrDecreasing && safeReport
}

func isReportSafeRevised(report []int) bool {
	if isReportSafe(report) {
		return true
	}

	possibleChanges := make([][]int, 0)

	for i := 0; i < len(report); i++ {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[:i])
		copy(newReport[i:], report[i+1:])

		possibleChanges = append(possibleChanges, newReport)
	}

	for _, possibleChange := range possibleChanges {
		if isReportSafe(possibleChange) {
			return true
		}
	}

	return false
}

type TransformDataPipe struct {
}

func (t *TransformDataPipe) Process(input common.PipelineContext) common.PipelineContext {
	reports := make([][]int, 0)
	safeReports := make([]bool, 0)
	safeReportsRevised := make([]bool, 0)

	callbackFinder := func(node *shared.ParseTree) (shared2.TransformCallback, int) {
		switch node.Symbol {
		case "report":
			return func(node *shared.ParseTree) {
				currentReport := make([]int, 0)

				for _, child := range node.Children {
					if child.Token.Type != shared3.NumberToken {
						panic(fmt.Sprintf("expected number token, got %v", child.Token.Type))
					}

					convertedValue, err := strconv.Atoi(string(child.Token.Value))

					if err != nil {
						panic(err)
					}

					currentReport = append(currentReport, convertedValue)
				}

				if isReportSafe(currentReport) {
					safeReports = append(safeReports, true)
				} else {
					safeReports = append(safeReports, false)
				}

				if isReportSafeRevised(currentReport) {
					safeReportsRevised = append(safeReportsRevised, true)
				} else {
					safeReportsRevised = append(safeReportsRevised, false)
				}

				reports = append(reports, currentReport)
			}, 0
		}
		return nil, 0
	}

	transformer := transforming.NewTransformer(
		callbackFinder,
	)
	transformer.Transform(input.ParseTree)

	safeReportNum := 0
	safeReportNumRevised := 0

	for _, safeReport := range safeReports {
		if safeReport {
			safeReportNum++
		}
	}

	for _, safeReport := range safeReportsRevised {
		if safeReport {
			safeReportNumRevised++
		}
	}

	input.Reports = reports
	input.SafeReports = safeReportNum
	input.SafeReportsRevised = safeReportNumRevised

	return input
}
