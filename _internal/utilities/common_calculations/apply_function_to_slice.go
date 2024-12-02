package common_calculations

// MapAndTransformSlice applies a given function to each element of a slice and stores the results.
//
// Parameters:
//   - slice: A pointer to the input slice of type []T.
//   - function: A function that takes a T and a map[T]int, and returns a U.
//   - dataMap: A map of type map[T]int used as additional input for the function.
//   - result: A pointer to the output slice of type []U where results will be stored.
//
// The function modifies the result slice in-place, populating it with the output
// of applying the given function to each element of the input slice.
func MapAndTransformSlice[T comparable, U any](slice *[]T, function func(T, map[T]int) U, dataMap map[T]int, result *[]U) {
	for i, item := range *slice {
		if i < len(*result) {
			(*result)[i] = function(item, dataMap)
		} else {
			*result = append(*result, function(item, dataMap))
		}
	}
}

func ApplyFunctionToSlice[T comparable](slice *[]T, function func(T) T) {
	for i, item := range *slice {
		if i < len(*slice) {
			(*slice)[i] = function(item)
		} else {
			*slice = append(*slice, function(item))
		}
	}
}
