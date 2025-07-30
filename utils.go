package utils

func InSlice(lines []string, pattern string) bool {
	for _, n := range lines {
		if pattern == n {
			return true
		}
	}
	return false
}

func InSliceInt(numbers []int, pattern int) bool {
	for _, n := range numbers {
		if pattern == n {
			return true
		}
	}
	return false
}
