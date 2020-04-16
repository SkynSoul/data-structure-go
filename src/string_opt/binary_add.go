package string_opt

import "math"

// 纯字符串比较
func AddBinary(a string, b string) string {
    aLen, bLen := len(a), len(b)
    if aLen == 0 {
        return a
    }
    if bLen == 0 {
        return b
    }
    ret := ""
    maxLen := int(math.Max(float64(aLen), float64(bLen)))
    aIdx, bIdx := aLen - 1, bLen - 1
    isCarry := false
    for i := 0; i < maxLen; i++ {
        var curA, curB uint8 = '0', '0'
        if aIdx >= 0 {
            curA = a[aIdx]
        }
        if bIdx >= 0 {
            curB = b[bIdx]
        }
        curRet := "0"
        if curA == '0' && curB == '0' {
            if isCarry {
                curRet = "1"
                isCarry = false
            }
        } else if curA == '1' && curB == '1' {
            if i == maxLen - 1 {
                if isCarry {
                    curRet = "11"
                } else {
                    curRet = "10"
                }
            } else {
                if isCarry {
                    curRet = "1"
                } else {
                    curRet = "0"
                }
            }
            isCarry = true
        } else {
            if i == maxLen - 1 {
                if isCarry {
                    curRet = "10"
                } else {
                    curRet = "1"
                }
            } else {
                if isCarry {
                    curRet = "0"
                } else {
                    curRet = "1"
                }
            }
        }
        ret = curRet + ret
        aIdx--
        bIdx--
    }
    return ret
}
