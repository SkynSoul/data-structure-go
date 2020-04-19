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

    nums = []int{1, 2, 3, 4, 5, 6, 7}
    k = 3
    t.Logf("nums is %v, k is %v", nums, k)
    Rotate(nums, k)
    t.Logf("after rotate2: %v", nums)

    nums = []int{-1, -100, 3, 99}
    k = 2
    t.Logf("nums is %v, k is %v", nums, k)
    Rotate2(nums, k)
    t.Logf("after rotate2: %v", nums)
}

func TestReverseWords(t *testing.T) {
    s := "the sky is blue"
    t.Logf("src string: %s, after reverse: %s", s, ReverseWords(s))
    s = "  hello world!  "
    t.Logf("src string: %s, after reverse: %s", s, ReverseWords(s))
    s = "a good   example"
    t.Logf("src string: %s, after reverse: %s", s, ReverseWords(s))
    s = "ab"
    t.Logf("src string: %s, after reverse: %s", s, ReverseWords(s))
}

func BenchmarkRotate(b *testing.B) {
    for i := 0; i < b.N; i++ {
        nums := []int{1, 2, 3, 4, 5, 6, 7}
        Rotate(nums, 3)
    }
}

func BenchmarkRotate2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        nums := []int{1, 2, 3, 4, 5, 6, 7}
        Rotate2(nums, 3)
    }
}