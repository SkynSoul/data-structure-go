package arr_string_opt

// "the sky is blue" -->  "blue is sky the"
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

// "Let's take LeetCode contest"  -->  "s'teL ekat edoCteeL tsetnoc"
func ReverseWords2(s string) string {
	byteArr := []byte(s)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			reverseString(byteArr[start : i])
			start = i + 1
		}
	}
	reverseString(byteArr[start:])
	return string(byteArr)
}

func reverseString(str []byte) {
	for left, right := 0, len(str) - 1; left < right; left, right = left + 1, right - 1 {
		str[left], str[right] = str[right], str[left]
	}
}