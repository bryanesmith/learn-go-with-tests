package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(arrays ...[]int) []int {

	// Alternatively, could do this:
	// 	lengthOfNumbers := len(numbersToSum)
	// 	sums := make([]int, lengthOfNumbers)

	var sums []int = []int{}

	for _, arr := range arrays {
		sums = append(sums, Sum(arr))
	}
	return sums
}

func SumAllTails(arrays ...[]int) []int {
	var sums []int = []int{}
	for _, arr := range arrays {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			tail := arr[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
