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