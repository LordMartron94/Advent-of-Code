package common_calculations

func SumInts(ints *[]int, sum *int) {
	for _, value := range *ints {
		*sum += value
	}
}

func SumIntsAndReturn(ints []int) int {
	sum := 0
	SumInts(&ints, &sum)
	return sum
}
