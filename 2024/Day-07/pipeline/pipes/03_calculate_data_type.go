package pipes

import (
	"context"
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-07/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-07/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/extensions"
)

type CalculateDataPipe struct{}

type Operand int

const (
	Plus Operand = iota
	Mult
	Combine
)

func (c *CalculateDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	result := 0
	resultPipe := 0

	for _, equation := range input.Equations {
		notPipe, pipe := c.validateEquationVariations(equation)

		if notPipe {
			result += equation.TestNumber
		}

		if pipe {
			resultPipe += equation.TestNumber
		}
	}

	input.Result = result
	input.ResultPipe = resultPipe
	return input
}

func (c *CalculateDataPipe) validateEquation(equation common.Equation, operands []Operand, combineAllowed bool) bool {
	if len(operands) != len(equation.Elements)-1 {
		panic("Invalid number of operands for equation")
	}

	resultOfEquation := equation.Elements[0]

	for i := range len(equation.Elements) - 1 {
		associatedOperand := operands[i]

		switch associatedOperand {
		case Plus:
			resultOfEquation += equation.Elements[i+1]
			continue
		case Mult:
			resultOfEquation *= equation.Elements[i+1]
			continue
		case Combine:
			if !combineAllowed {
				continue
			}
			num, err := strconv.Atoi(strconv.Itoa(resultOfEquation) + strconv.Itoa(equation.Elements[i+1]))

			if err != nil {
				panic(err)
			}

			resultOfEquation = num
		}
	}

	if resultOfEquation == equation.TestNumber {
		return true
	}

	return false
}

func (c *CalculateDataPipe) validateEquationVariations(equation common.Equation) (bool, bool) {
	standardOperands := make([]Operand, len(equation.Elements)-1)

	for i := range standardOperands {
		standardOperands[i] = Plus
	}

	return c.checkBranchesPipeNot(standardOperands, equation), c.checkBranchesPipe(standardOperands, equation)
}

func (c *CalculateDataPipe) checkBranchesPipeNot(operands []Operand, equation common.Equation) bool {
	ctx, cancel := context.WithCancel(context.Background())

	sequenceCorrect := false

	extensions.ApplyFunctionToGeneratedBinaryVariationsGeneric(operands, Plus, Mult, func(a, b Operand) bool {
		return a == b
	}, func(slice []Operand) {
		if c.validateEquation(equation, slice, false) {
			cancel()
			sequenceCorrect = true
		}
	}, ctx)

	return sequenceCorrect
}

func (c *CalculateDataPipe) checkBranchesPipe(operands []Operand, equation common.Equation) bool {
	ctx, cancel := context.WithCancel(context.Background())

	sequenceCorrect := false

	extensions.ApplyFunctionToGeneratedVariationsGeneric(operands, Plus, func(operand Operand) Operand {
		switch operand {
		case Plus:
			return Mult
		case Mult:
			return Combine
		case Combine:
			return Plus
		default:
			panic("Invalid operand type")
		}
	}, []Operand{Plus, Mult, Combine}, func(a, b Operand) bool {
		return a == b
	}, func(slice []Operand) {
		if c.validateEquation(equation, slice, true) {
			cancel()
			sequenceCorrect = true
		}
	}, ctx)

	return sequenceCorrect
}
