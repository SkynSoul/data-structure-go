package arr_string_opt

func ReverseWords(s string) string {
	stackHelper := make([]string, 0)
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(word) > 0 {
				stackHelper = append(stackHelper, word)
				word = ""
			}
			continue
		}
		word += string(s[i])
	}
	if len(word) > 0 {
		stackHelper = append(stackHelper, word)
	}

	if len(stackHelper) == 0 {
		return ""
	}
	ret := ""
	for i := len(stackHelper) - 1; i > 0; i-- {
		ret += stackHelper[i]
		ret += " "
	}
	ret += stackHelper[0]
	return ret
}