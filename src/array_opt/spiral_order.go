package array_opt

func SpiralOrder(matrix [][]int) []int {
    xMax := len(matrix)
    if xMax == 0 {
        return []int{}
    }
    yMax := len(matrix[0])
    totalNum := xMax * yMax
    ret := make([]int, 0, totalNum)
    x, y, dir, step := 0, 0, 0, 0
    for len(ret) < totalNum {
        ret = append(ret, matrix[x][y])
        step++
        switch dir {
        case 0: // 向右
            if step < yMax {
                y += 1
            } else {    // 到达边界  向下
                dir = (dir + 1) % 4
                step = 0
                xMax--
                x += 1
            }
            break
        case 1: // 向下
            if step < xMax {
                x += 1
            } else {    // 到达边界  向左
                dir = (dir + 1) % 4
                step = 0
                yMax--
                y -= 1
            }
            break
        case 2: // 向左
            if step < yMax {
                y -= 1
            } else {    // 到达边界  向上
                dir = (dir + 1) % 4
                step = 0
                xMax--
                x -= 1
            }
            break
        case 3: // 向上
            if step < xMax {
                x -= 1
            } else {    // 到达边界  向右
                dir = (dir + 1) % 4
                step = 0
                yMax--
                y += 1
            }
            break
        default:
            return ret
        }
    }
    return ret
}