package queue_stack

import "strconv"

// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数
func DecodeString(s string) string {
    stackIn := make([]string, 0, 10)
    for i := 0; i < len(s); i++ {
        if s[i] == ']' {
            str, numStr, curRet := "", "", ""
            readNum := false
            for len(stackIn) > 0 {
                cur := stackIn[len(stackIn) - 1]
                stackIn = stackIn[:len(stackIn) - 1]
                if cur == "[" && !readNum {
                    readNum = true
                    continue
                }
                if !readNum {
                    str = cur + str
                    continue
                }
                if !isNum(cur) {
                    stackIn = append(stackIn, cur)
                    break
                }
                numStr = cur + numStr
            }
            num, _ := strconv.ParseInt(numStr, 10, 64)
            for j := 0; j < int(num); j++ {
                curRet += str
            }
            stackIn = append(stackIn, curRet)
            continue
        }
        stackIn = append(stackIn, string(s[i]))
    }
    ret := ""
    for len(stackIn) > 0 {
        ret = stackIn[len(stackIn) - 1] + ret
        stackIn = stackIn[:len(stackIn) - 1]
    }
    return ret
}

func isNum(str string) bool {
    byteArr := []byte(str)
    if len(byteArr) == 1 && byteArr[0] <= '9' && byteArr[0] >= '0' {
        return true
    }
    return false
}
