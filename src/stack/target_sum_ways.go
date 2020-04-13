package stack

import (
    "container/list"
)

type SumNode struct {
    Val int
    Depth int
}

// 深度优先穷举算法  无优化  N为数组长度 时间复杂度O(2^N) 空间复杂度O(N)
func FindTargetSumWays(nums []int, S int) int {
    ways, numLen := 0, len(nums)
    stackHelper := list.New()
    stackHelper.PushBack(&SumNode{Val: 0, Depth: 0})
    for stackHelper.Len() > 0 {
        curNode := stackHelper.Back().Value.(*SumNode)
        if curNode.Val == S && curNode.Depth == numLen {
            ways++
        }
        stackHelper.Remove(stackHelper.Back())
        if curNode.Depth < numLen {
            stackHelper.PushBack(&SumNode{Val: curNode.Val + nums[curNode.Depth], Depth: curNode.Depth + 1})
            stackHelper.PushBack(&SumNode{Val: curNode.Val - nums[curNode.Depth], Depth: curNode.Depth + 1})
        }
    }
    return ways
}

func FindTargetSumWays2(nums []int, S int) int {
    ways, numLen := 0, len(nums)
    stackHelper := make([]*SumNode, 0)
    stackHelper = append(stackHelper, &SumNode{Val: 0, Depth: 0})
    for len(stackHelper) > 0 {
        curNode := stackHelper[len(stackHelper) - 1]
        if curNode.Val == S && curNode.Depth == numLen {
            ways++
        }
        stackHelper = stackHelper[:len(stackHelper) - 1]
        if curNode.Depth < numLen {
            stackHelper = append(stackHelper, &SumNode{Val: curNode.Val + nums[curNode.Depth], Depth: curNode.Depth + 1})
            stackHelper = append(stackHelper, &SumNode{Val: curNode.Val - nums[curNode.Depth], Depth: curNode.Depth + 1})
        }
    }
    return ways
}

func FindTargetSumWays3(nums []int, S int) int {
    return dfsSum(nums, S, 0, 0)
}

func dfsSum(nums []int, target int, sum int, depth int) int {
    ways := 0
    if depth == len(nums) && sum == target {
        ways++
    }
    if depth < len(nums) {
        ways += dfsSum(nums, target, sum + nums[depth], depth + 1)
        ways += dfsSum(nums, target, sum - nums[depth], depth + 1)
    }
    return ways
}

// 动态规划解法  N为数组长度，sum为和的可能个数  时间复杂度O(N*sum) 空间复杂度O(N*sum)
// 一个大问题拆成若干小问题，且***这些小问题能够被重复复用***
func FindTargetSumWays4(nums []int, S int) int {
    // 限定条件：1、数组长度不超过20 2、所有元素绝对值之和不超过1000
    // 最终结果范围[-1000, 1000]，共2001种可能（包括0）
    // 深度为i + 1（数组从零开始），结果为j，代表的可能路径个数有如下表达式
    // dp[i][j] = dp[i - 1][j - nums[i]] + dp[i - 1][j + nums[i]]
    // 转换如下
    // dp[i][j] = 0
    // dp[i][j] += dp[i - 1][j - nums[i]]
    // dp[i][j] += dp[i - 1][j + nums[i]]
    // j的下标不为负，统一加上1000，则j的取值范围可规范为[0, 2000]
    if S > 1000 {
        return 0
    }
    dp := [20][2001]int{}
    // 初始深度为1的情况
    dp[0][0 + nums[0] + 1000] += 1
    dp[0][0 - nums[0] + 1000] += 1
    for i := 1; i < len(nums); i++ {
        for sum := -1000; sum <= 1000; sum++ {
            if sum - nums[i] + 1000 >= 0 {
               dp[i][sum + 1000] += dp[i - 1][sum - nums[i] + 1000]
            }
            if sum + nums[i] + 1000 <= 2000 {
               dp[i][sum + 1000] += dp[i - 1][sum + nums[i] + 1000]
            }
        }
    }
    return dp[len(nums) - 1][S + 1000]
}