package utils

func Contains(line []string, pattern string) bool {
	for _, n := range line {
		if pattern == n {
			return true
		}
	}
	return false
}
