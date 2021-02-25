package iteration

func Repeat(character string, times int) string {
	var repeatCount = times
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
