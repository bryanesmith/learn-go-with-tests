package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	var sums []int = []int{}

	for _, arr := range arrays {
		sums = append(sums, Sum(arr))
	}

	return sums
}
