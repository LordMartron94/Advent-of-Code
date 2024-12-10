package extensions

import (
	"context"
	"fmt"
	"slices"
)

// FindNumberOfMatchesInSlice returns the amount of occurrences for the target value in the given slice.
// If reverseAllowed is true, the search will be performed in reverse order as well.
func FindNumberOfMatchesInSlice[T comparable](slice []T, targets []T, reverseAllowed bool) int {
	matches := 0

	for i := 0; i < len(slice)-len(targets)+1; i++ {
		match := true
		for j := 0; j < len(targets); j++ {
			if slice[i+j] != targets[j] {
				match = false
				break
			}
		}
		if match {
			matches++
		}
	}

	if reverseAllowed {
		newTargets := make([]T, len(targets))
		copy(newTargets, targets)
		slices.Reverse(newTargets)

		matches += FindNumberOfMatchesInSlice[T](slice, targets, false)
	}

	return matches
}

func FindNumberOfMatchesInSliceV2[T any](slice []T, targets []T, reverseAllowed bool, equalityCheckFunc func(a, b T) bool) int {
	matches := 0

	for i := 0; i < len(slice)-len(targets)+1; i++ {
		match := true
		for j := 0; j < len(targets); j++ {
			if !equalityCheckFunc(slice[i+j], targets[j]) {
				match = false
				break
			}
		}
		if match {
			matches++
		}
	}

	if reverseAllowed {
		newTargets := make([]T, len(targets))
		copy(newTargets, targets)
		slices.Reverse(newTargets)

		matches += FindNumberOfMatchesInSliceV2[T](slice, targets, false, equalityCheckFunc)
	}

	return matches
}

// GetFormattedString returns a formatted string representation of the given slice.
func GetFormattedString[T any](slice []T) string {
	formattedString := ""
	for _, item := range slice {
		formattedString += fmt.Sprintf("%v, ", item)
	}
	if len(formattedString) > 0 {
		formattedString = "[" + formattedString[:len(formattedString)-2] + "]"
	} else {
		formattedString = "[" + formattedString + "]"
	}

	return formattedString
}

// GetFormattedStringNil returns a formatted string representation of the given slice.
func GetFormattedStringNil[T any](slice []*T) string {
	formattedString := ""
	for _, item := range slice {
		if item == nil {
			formattedString += "nil, "
		} else {
			formattedString += fmt.Sprintf("%v, ", *item)
		}
	}
	if len(formattedString) > 0 {
		formattedString = "[" + formattedString[:len(formattedString)-2] + "]"
	} else {
		formattedString = "[" + formattedString + "]"
	}

	return formattedString
}

func switchBinary[T comparable](item, from, to T) T {
	if item == from {
		return to
	} else if item == to {
		return from
	}

	panic("Invalid binary switch")
}

func validateBinarySequence[T comparable](slice []T, from T) bool {
	for _, item := range slice {
		if item != from {
			return false
		}
	}
	return true
}

// GenerateBinaryVariations generates all possible variations of a given binary value sequence.
// Requires all elements in the slice to start as 'from'
func GenerateBinaryVariations[T comparable](slice []T, from, to T) [][]T {
	if !validateBinarySequence(slice, from) {
		panic("All elements in the slice should start as 'from'")
	}

	n := len(slice)

	variations := make([][]T, 2<<(n-1))
	variations[0] = slice

	for v := 1; v < len(variations); v++ {
		variation := make([]T, n)
		copy(variation, variations[v-1])

		for i := len(variation) - 1; i >= 0; i-- {
			if variation[i] == from {
				variation[i] = switchBinary(variation[i], from, to)
				break
			} else {
				variation[i] = switchBinary(variation[i], from, to)
			}
		}

		variations[v] = variation
	}

	if len(variations) != 2<<(n-1) {
		panic(fmt.Sprintf("Expected 2^%d variations, but got %d", n, len(variations)))
	}

	return variations
}

func switchBinaryV2[T any](item, from, to T, equalityFunc func(a, b T) bool) T {
	if equalityFunc(item, from) {
		return to
	} else if equalityFunc(item, to) {
		return from
	}

	panic("Invalid binary switch")
}

func validateSequence[T any](slice []T, from T, equalityFunc func(a, b T) bool) bool {
	for _, item := range slice {
		if !equalityFunc(item, from) {
			return false
		}
	}
	return true
}

// GenerateBinaryVariationsGeneric generates all possible variations of a given binary value sequence.
// Requires all elements in the slice to start as 'from'
// This version uses a custom equality function to compare items in the slice.
func GenerateBinaryVariationsGeneric[T any](slice []T, from, to T, equalityFunc func(a, b T) bool) [][]T {
	if !validateSequence(slice, from, equalityFunc) {
		panic("All elements in the slice should start as 'from'")
	}

	n := len(slice)

	variations := make([][]T, 2<<(n-1))
	variations[0] = slice

	for v := 1; v < len(variations); v++ {
		variation := make([]T, n)
		copy(variation, variations[v-1])

		for i := len(variation) - 1; i >= 0; i-- {
			if equalityFunc(variation[i], from) {
				variation[i] = switchBinaryV2(variation[i], from, to, equalityFunc)
				break
			} else {
				variation[i] = switchBinaryV2(variation[i], from, to, equalityFunc)
			}
		}

		variations[v] = variation
	}

	if len(variations) != 2<<(n-1) {
		panic(fmt.Sprintf("Expected 2^%d variations, but got %d", n, len(variations)))
	}

	return variations
}

// ApplyFunctionToGeneratedBinaryVariationsGeneric generates all possible variations of a given binary value sequence and applies function to them.
// Requires all elements in the slice to start as 'from'
// This version uses a custom equality function to compare items in the slice.
func ApplyFunctionToGeneratedBinaryVariationsGeneric[T any](slice []T, from, to T, equalityFunc func(a, b T) bool, applyFunc func(slice []T), ctx context.Context) {
	if !validateSequence(slice, from, equalityFunc) {
		panic("All elements in the slice should start as 'from'")
	}

	applyFunc(slice)

	n := len(slice)

	variations := make([][]T, 2<<(n-1))
	variations[0] = slice

	for v := 1; v < len(variations); v++ {
		select {
		case <-ctx.Done():
			return
		default:
			variation := make([]T, n)
			copy(variation, variations[v-1])

			for i := len(variation) - 1; i >= 0; i-- {
				variation[i] = switchBinaryV2(variation[i], from, to, equalityFunc)
				if equalityFunc(variation[i], from) {
					break
				}
			}

			applyFunc(variation)
			variations[v] = variation
		}
	}

	if len(variations) != 2<<(n-1) {
		panic(fmt.Sprintf("Expected 2^%d variations, but got %d", n, len(variations)))
	}
}

func IntPow(n, m int) int {
	result := 1
	for i := 0; i < m; i++ {
		result *= n
	}
	return result
}

// ApplyFunctionToGeneratedVariationsGeneric generates all possible variations of a given sequence
// by repeatedly applying a "next state" function to each element until a base state is reached.
// Requires all elements in the slice to start as 'from'
func ApplyFunctionToGeneratedVariationsGeneric[T any](slice []T, from T, nextState func(T) T, stateOptions []T, equalityFunc func(a, b T) bool, applyFunc func(slice []T), ctx context.Context) {
	if !validateSequence(slice, from, equalityFunc) {
		panic("All elements in the slice should start as 'from'")
	}

	applyFunc(slice)

	n := len(slice)
	totalVariationsExpected := IntPow(len(stateOptions), n)

	variations := make([][]T, totalVariationsExpected)
	variations[0] = slice

	for v := 1; v < len(variations); v++ {
		select {
		case <-ctx.Done():
			return
		default:
			variation := make([]T, n)
			copy(variation, variations[v-1])

			for i := len(variation) - 1; i >= 0; i-- {
				variation[i] = nextState(variation[i])
				if !equalityFunc(variation[i], from) {
					break
				}
			}

			applyFunc(variation)
			variations[v] = variation
		}
	}

	if len(variations) != totalVariationsExpected {
		panic(fmt.Sprintf("Expected %d variations, but got %d", totalVariationsExpected, len(variations)))
	}
}

func SliceContainsEqualityComparer[T any](slice []T, item T, equalityFunc func(a, b T) bool) bool {
	for _, v := range slice {
		if equalityFunc(v, item) {
			return true
		}
	}
	return false
}

func SliceGetIndexOfEqualityComparer[T any](slice []T, item T, equalityFunc func(a, b T) bool) int {
	for i, v := range slice {
		if equalityFunc(v, item) {
			return i
		}
	}
	return -1
}
