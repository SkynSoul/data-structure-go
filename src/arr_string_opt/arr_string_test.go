package arr_string_opt

import "testing"

func TestRotate(t *testing.T) {
    nums := []int{1, 2, 3, 4, 5, 6, 7}
    k := 3
    t.Logf("nums is %v, k is %v", nums, k)
    Rotate(nums, k)
    t.Logf("after rotate: %v", nums)
    
    nums = []int{-1, -100, 3, 99}
    k = 2
    t.Logf("nums is %v, k is %v", nums, k)
    Rotate(nums, k)
    t.Logf("after rotate: %v", nums)
}