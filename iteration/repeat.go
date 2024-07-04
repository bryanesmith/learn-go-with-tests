package iteration

func Repeat(input string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += input
	}
	return repeated
}
