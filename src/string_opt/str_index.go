package string_opt

func StrStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		start := i
		isFind := true
		for j := 0; j < len(needle); j++ {
			if start >= len(haystack) {
				return -1
			}
			if haystack[start] != needle[j] {
				isFind = false
				break
			}
			start++
		}
		if isFind {
			return i
		}
	}
	return -1
}
