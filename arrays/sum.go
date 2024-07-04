package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	cnt := len(arrays)
	sums := make([]int, cnt)

	for i, arr := range arrays {
		sums[i] = Sum(arr)
	}

	return sums
}
