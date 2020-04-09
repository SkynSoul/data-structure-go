package stack

import "container/list"

func IsValid(s string) bool {
    stackHelper := list.New()
    mapHelper := map[byte]byte{')': '(', ']': '[', '}': '{'}
    for _, c := range s {
        switch c {
        case '(', '[', '{':
            stackHelper.PushBack(byte(c))
            break
        case ')', ']', '}':
            if stackHelper.Len() == 0 || stackHelper.Back().Value.(byte) != mapHelper[byte(c)] {
               return false
            }
            stackHelper.Remove(stackHelper.Back())
        default:
            return false
        }
    }
    if stackHelper.Len() != 0 {
        return false
    }
    return true
}
