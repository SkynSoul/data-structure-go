package stack

import (
    "container/list"
    "reflect"
    "strconv"
)

func EvalRPNDesc(tokens []string) int {
    stackHelper := list.New()
    // 默认给定逆波兰表达式为合法的，这里不做异常处理
    for i := len(tokens) - 1; i >= 0; i-- {
        switch tokens[i] {
        case "+", "-", "*", "/":
            stackHelper.PushBack(tokens[i])
            break
        default:
            num, _ := strconv.ParseInt(tokens[i], 10, 64)
            for stackHelper.Len() > 0 && reflect.TypeOf(stackHelper.Back().Value).Name() == "int" {
                optNum := stackHelper.Back().Value.(int)
                stackHelper.Remove(stackHelper.Back())
                opt := stackHelper.Back().Value.(string)
                stackHelper.Remove(stackHelper.Back())
                switch opt {
                case "+":
                    num += int64(optNum)
                    break
                case "-":
                    num -= int64(optNum)
                    break
                case "*":
                    num *= int64(optNum)
                    break
                case "/":
                    num /= int64(optNum)
                    break
                default:
                    break
                }
            }
            stackHelper.PushBack(int(num))
            break
        }
    }
    return stackHelper.Back().Value.(int)
}

func EvalRPNAsc(tokens []string) int {
    a, b := 0, 0
    stackHelper := list.New()
    // 默认给定逆波兰表达式为合法的，这里不做异常处理
    for i := 0; i < len(tokens); i++ {
        switch tokens[i] {
        case "+":
            a, b = getOptNum(stackHelper)
            stackHelper.PushBack(b + a)
            break
        case "-":
            a, b = getOptNum(stackHelper)
            stackHelper.PushBack(b - a)
            break
        case "*":
            a, b = getOptNum(stackHelper)
            stackHelper.PushBack(b * a)
            break
        case "/":
            a, b = getOptNum(stackHelper)
            stackHelper.PushBack(b / a)
            break
        default:
            num, _ := strconv.ParseInt(tokens[i], 10, 64)
            stackHelper.PushBack(int(num))
            break
        }
    }
    return stackHelper.Back().Value.(int)
}

func getOptNum(stackHelper *list.List) (int, int) {
    a := stackHelper.Back().Value.(int)
    stackHelper.Remove(stackHelper.Back())
    b := stackHelper.Back().Value.(int)
    stackHelper.Remove(stackHelper.Back())
    return a, b
}