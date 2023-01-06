package repeater

func repeater(toRepeat string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += toRepeat
	}
	return repeated
}
