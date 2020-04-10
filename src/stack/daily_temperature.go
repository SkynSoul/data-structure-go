package stack

func DailyTemperatures(T []int) []int {
    srcLen := len(T)
    maxStack := make([]int, srcLen)
    for i := srcLen - 1; i >= 0; i-- {
       curVal := T[i]
       if i < srcLen - 1 && maxStack[i + 1] > curVal {
           curVal = maxStack[i + 1]
       }
        maxStack[i] = curVal
    }
    
    dstArr := make([]int, srcLen)
    for i := 0; i < srcLen; i++ {
        if maxStack[i] <= T[i] {
            dstArr[i] = 0
            continue
        }
        tarIdx := i + 1
        for j := i + 1; j < srcLen; j++ {
            if T[i] < T[j] {
                tarIdx = j
                break
            }
        }
        dstArr[i] = tarIdx - i
    }
    return dstArr
}

func DailyTempWithoutStack(T []int) []int {
    srcLen := len(T)
    dstArr := make([]int, srcLen)
    for i := 0; i < srcLen; i++ {
        tarIdx := i + 1
        isFind := false
        for j := i + 1; j < srcLen; j++ {
            if T[i] < T[j] {
                tarIdx = j
                isFind = true
                break
            }
        }
        if isFind {
            dstArr[i] = tarIdx - i
        } else {
            dstArr[i] = 0
        }
    }
    return dstArr
}