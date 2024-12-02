package common_calculations

func GetPairs(slice1, slice2 []int) [][2]int {
	if len(slice1) != len(slice2) {
		panic("Slices must have the same length")
	}

	pairs := make([][2]int, len(slice1))

	for i := range slice1 {
		pairs[i] = [2]int{slice1[i], slice2[i]}
	}

	return pairs
}
