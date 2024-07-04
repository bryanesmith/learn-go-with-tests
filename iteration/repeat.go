package iteration

func Repeat(input string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += input
	}
	return repeated
}
