package hash_table

import "testing"

func TestHashSet(t *testing.T) {
    mySet := ConstructorHashSet()
    t.Logf("mySet.Add(1)")
    mySet.Add(1)
    t.Logf("mySet.Add(2)")
    mySet.Add(2)
    t.Logf("mySet.Contains(1): %t", mySet.Contains(1))
    t.Logf("mySet.Contains(2): %t", mySet.Contains(2))
    t.Logf("mySet.Contains(3): %t", mySet.Contains(3))
    t.Logf("mySet.Add(2)")
    mySet.Add(2)
    t.Logf("mySet.Contains(2): %t", mySet.Contains(2))
    t.Logf("mySet.Remove(2)")
    mySet.Remove(2)
    t.Logf("mySet.Contains(2): %t", mySet.Contains(2))
    t.Logf("mySet.Remove(2)")
    mySet.Remove(2)
    t.Logf("mySet.Remove(10002)")
    mySet.Remove(10002)
    t.Logf("mySet.Contains(10002): %t", mySet.Contains(10002))
    t.Logf("mySet.Add(2)")
    mySet.Add(2)
    t.Logf("mySet.Add(10002)")
    mySet.Add(10002)
    t.Logf("mySet.Contains(10002): %t", mySet.Contains(10002))
    t.Logf("mySet.Remove(10002)")
    mySet.Remove(10002)
}

func TestHashMap(t *testing.T) {
    myMap := ConstructorHashMap()
    t.Logf("myMap.Put(1, 1)")
    myMap.Put(1, 1)
    t.Logf("myMap.Put(2, 2)")
    myMap.Put(2, 2)
    t.Logf("myMap.Get(1): %d", myMap.Get(1))
    t.Logf("myMap.Get(2): %d", myMap.Get(2))
    t.Logf("myMap.Get(3): %d", myMap.Get(3))
    t.Logf("myMap.Put(2, 100)")
    myMap.Put(2, 100)
    t.Logf("myMap.Get(2): %d", myMap.Get(2))
    t.Logf("myMap.Remove(2)")
    myMap.Remove(2)
    t.Logf("myMap.Get(2): %d", myMap.Get(2))
}

func TestContainsDuplicate(t *testing.T) {
    nums := []int{1, 2, 3, 1}
    t.Logf("the nums is %v, is contains duplicate: %t", nums, ContainsDuplicate(nums))
    nums = []int{1, 2, 3, 4}
    t.Logf("the nums is %v, is contains duplicate: %t", nums, ContainsDuplicate(nums))
}

func TestSingleNumber(t *testing.T) {
    nums := []int{1, 2, 1}
    t.Logf("the nums is %v, single number: %d", nums, SingleNumber(nums))
    nums = []int{4, 1, 2, 1, 2}
    t.Logf("the nums is %v, single number: %d", nums, SingleNumber(nums))
}

func TestSingleNumberXOR(t *testing.T) {
    nums := []int{1, 2, 1}
    t.Logf("the nums is %v, single number: %d", nums, SingleNumberXOR(nums))
    nums = []int{4, 1, 2, 1, 2}
    t.Logf("the nums is %v, single number: %d", nums, SingleNumberXOR(nums))
}

func TestIntersection(t *testing.T) {
    nums1 := []int{1, 2, 2, 1}
    nums2 := []int{2, 2}
    t.Logf("nums1: %v, nums2: %v, intersection: %v", nums1, nums2, Intersection(nums1, nums2))
}

func TestIsHappy(t *testing.T) {
    num := 19
    t.Logf("number is %d, is happy: %t", num, IsHappy(num))
    num = 100
    t.Logf("number is %d, is happy: %t", num, IsHappy(num))
    num = 123
    t.Logf("number is %d, is happy: %t", num, IsHappy(num))
    num = 496465
    t.Logf("number is %d, is happy: %t", num, IsHappy(num))
}

func TestTwoSum(t *testing.T) {
    nums := []int{2, 7, 11, 15}
    target := 9
    t.Logf("nums: %v, target: %v, ret: %v", nums, target, TwoSum(nums, target))
    nums = []int{2, 5, 5, 15}
    target = 10
    t.Logf("nums: %v, target: %v, ret: %v", nums, target, TwoSum(nums, target))
    nums = []int{3, 2, 4, 15}
    target = 6
    t.Logf("nums: %v, target: %v, ret: %v", nums, target, TwoSum(nums, target))
    nums = []int{1, 3, 4, 2}
    target = 6
    t.Logf("nums: %v, target: %v, ret: %v", nums, target, TwoSum(nums, target))
}

func TestIsIsomorphic(t *testing.T) {
    str1, str2 := "egg", "add"
    t.Logf("s: %s, t: %s, isomorphic: %t", str1, str2, IsIsomorphic2(str1, str2))
    str1, str2 = "foo", "bar"
    t.Logf("s: %s, t: %s, isomorphic: %t", str1, str2, IsIsomorphic2(str1, str2))
    str1, str2 = "paper", "title"
    t.Logf("s: %s, t: %s, isomorphic: %t", str1, str2, IsIsomorphic2(str1, str2))
}

func BenchmarkSingleNumber(b *testing.B) {
    nums := []int{4, 1, 2, 1, 2}
    for i := 0; i < b.N; i++ {
        SingleNumber(nums)
    }
}

func BenchmarkSingleNumberXOR(b *testing.B) {
    nums := []int{4, 1, 2, 1, 2}
    for i := 0; i < b.N; i++ {
        SingleNumberXOR(nums)
    }
}

func BenchmarkIsIsomorphic(b *testing.B) {
    str1, str2 := "paper", "title"
    for i := 0; i < b.N; i++ {
        IsIsomorphic(str1, str2)
    }
}

func BenchmarkIsIsomorphic2(b *testing.B) {
    str1, str2 := "paper", "title"
    for i := 0; i < b.N; i++ {
        IsIsomorphic2(str1, str2)
    }
}