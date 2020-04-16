package string_opt

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	srcStr := strs[0]
	i := 0
	for ; i < len(srcStr); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || srcStr[i] != strs[j][i] {
				return srcStr[:i]
			}
		}
	}
	return srcStr[:i]
}