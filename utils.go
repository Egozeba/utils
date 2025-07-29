package utils

func Contains(lines []string, pattern string) bool {
	for _, n := range lines {
		if pattern == n {
			return true
		}
	}
	return false
}

func ContainsInt(numbers []int, pattern int) bool {
	for _, n := range numbers {
		if pattern == n {
			return true
		}
	}
	return false
}
