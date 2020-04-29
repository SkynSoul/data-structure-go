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

func TestGolangMap(t *testing.T) {
    testMap := make(map[int]map[int]bool)
    t.Logf("%v", testMap[1])
    if testMap[2] == nil {
        t.Logf("111111111111")
    } else {
        t.Logf("222222222222")
    }
    testMap[1] = make(map[int]bool)
    target := testMap[1]
    target[123] = true
    t.Logf("%v", testMap)
    target = testMap[2]
    //target[456] = true
    //t.Logf("%v", testMap)
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

func TestFindRestaurant(t *testing.T) {
   list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
   list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
   t.Logf("find restaurant: %v", FindRestaurant(list1, list2))
   list1 = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
   list2 = []string{"KFC", "Shogun", "Burger King"}
   t.Logf("find restaurant: %v", FindRestaurant(list1, list2))
}

func TestFirstUniqChar(t *testing.T) {
    s := "leetcode"
    t.Logf("s: %s, first unique char idx: %d", s, FirstUniqChar(s))
    s = "loveleetcode"
    t.Logf("s: %s, first unique char idx: %d", s, FirstUniqChar(s))
    s = "ee"
    t.Logf("s: %s, first unique char idx: %d", s, FirstUniqChar(s))
    s = "ea"
    t.Logf("s: %s, first unique char idx: %d", s, FirstUniqChar(s))
    s = ""
    t.Logf("s: %s, first unique char idx: %d", s, FirstUniqChar(s))
    s = "a"
    t.Logf("s: %s, first unique char idx: %d", s, FirstUniqChar(s))
}

func TestIntersect(t *testing.T) {
    nums1 := []int{1, 2, 2, 1}
    nums2 := []int{2, 2}
    t.Logf("the intersect is %v", Intersect(nums1, nums2))
    nums1 = []int{4, 9, 5}
    nums2 = []int{9, 4, 9, 8, 4}
    t.Logf("the intersect is %v", Intersect(nums1, nums2))
}

func TestContainsNearbyDuplicate(t *testing.T) {
    nums := []int{1, 2, 3, 1}
    k := 3
    t.Logf("nums: %v, k: %d, ret: %t", nums, k, ContainsNearbyDuplicate(nums, k))
    nums = []int{1, 0, 1, 1}
    k = 1
    t.Logf("nums: %v, k: %d, ret: %t", nums, k, ContainsNearbyDuplicate(nums, k))
    nums = []int{1, 2, 3, 1, 2, 3}
    k = 2
    t.Logf("nums: %v, k: %d, ret: %t", nums, k, ContainsNearbyDuplicate(nums, k))
}

func TestGroupAnagrams(t *testing.T) {
    strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    t.Logf("ret: %v", GroupAnagrams(strs))
}

func TestIsValidSudoku(t *testing.T) {
    board := [][]byte{
        {'5', '3' , '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.' , '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9' , '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.' , '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.' , '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.' , '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6' , '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.' , '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.' , '.', '.', '8', '.', '.', '7', '9'},
    }
    t.Logf("is sudo: %t", IsValidSudoku(board))
    t.Logf("is sudo: %t", IsValidSudokuByArr(board))
}

func TestFindDuplicateSubtrees(t *testing.T) {
    n7 := &TreeNode{Val: 4}
    n6 := &TreeNode{Val: 4}
    n5 := &TreeNode{Val: 2, Left: n6}
    n4 := &TreeNode{Val: 3, Left: n5, Right: n7}
    n3 := &TreeNode{Val: 4}
    n2 := &TreeNode{Val: 2, Left: n3}
    n1 := &TreeNode{Val: 1, Left: n2, Right: n4}
    retArr := FindDuplicateSubtrees(n1)
    for _, val := range retArr {
        t.Logf("cur root, addr: %p, val: %v, preorder ret: %s", val, val, preorder(val, make(map[string][]*TreeNode)))
    }
}

func TestNumJewelsInStones(t *testing.T) {
    J, S := "aA", "aAAbbbb"
    t.Logf("J: %s, S: %s, num: %d", J, S, NumJewelsInStones(J, S))
    J, S = "z", "ZZ"
    t.Logf("J: %s, S: %s, num: %d", J, S, NumJewelsInStones(J, S))

    J, S = "aA", "aAAbbbb"
    t.Logf("J: %s, S: %s, num: %d", J, S, NumJewelsInStonesByArr(J, S))
    J, S = "z", "ZZ"
    t.Logf("J: %s, S: %s, num: %d", J, S, NumJewelsInStonesByArr(J, S))
}

func TestLengthOfLongestSubstring(t *testing.T) {
    s := "abcabcbb"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring(s))
    s = "bbbbbbbbbbb"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring(s))
    s = "pwwkew"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring(s))
    s = " "
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring(s))
    s = "bpfbhmipx"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring(s))
    s = "umvejcuuk"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring(s))
}

func TestLengthOfLongestSubstring2(t *testing.T) {
    s := "abcabcbb"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring2(s))
    s = "bbbbbbbbbbb"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring2(s))
    s = "pwwkew"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring2(s))
    s = " "
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring2(s))
    s = "bpfbhmipx"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring2(s))
    s = "umvejcuuk"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring2(s))
}

func TestLengthOfLongestSubstring3(t *testing.T) {
    s := "abcabcbb"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring3(s))
    s = "bbbbbbbbbbb"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring3(s))
    s = "pwwkew"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring3(s))
    s = " "
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring3(s))
    s = "bpfbhmipx"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring3(s))
    s = "umvejcuuk"
    t.Logf("the length of longest sub string is %d", LengthOfLongestSubstring3(s))
}

func TestFourSumCount(t *testing.T) {
   A := []int{1, 2}
   B := []int{-2, -1}
   C := []int{-1, 2}
   D := []int{0, 2}
   t.Logf("four sum count is %d", FourSumCount(A, B, C, D))
}

func TestTopKFrequent(t *testing.T) {
    nums := []int{1, 1, 1, 2, 2, 3}
    k := 2
    t.Logf("nums: %v, k: %d, the ret: %v", nums, k, TopKFrequent(nums, k))
    nums = []int{1}
    k = 1
    t.Logf("nums: %v, k: %d, the ret: %v", nums, k, TopKFrequent(nums, k))
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

func BenchmarkIsValidSudoku(b *testing.B) {
    board := [][]byte{
        {'5', '3' , '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.' , '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9' , '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.' , '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.' , '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.' , '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6' , '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.' , '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.' , '.', '.', '8', '.', '.', '7', '9'},
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        IsValidSudoku(board)
    }
}

func BenchmarkIsValidSudokuByArr(b *testing.B) {
    board := [][]byte{
        {'5', '3' , '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.' , '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9' , '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.' , '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.' , '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.' , '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6' , '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.' , '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.' , '.', '.', '8', '.', '.', '7', '9'},
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        IsValidSudokuByArr(board)
    }
}

func BenchmarkNumJewelsInStones(b *testing.B) {
    J, S := "aAweio", "aAAbbbbaapawerpaAFWEFAWJOAWRIErwearkaweapwej"
    for i := 0; i < b.N; i++ {
        NumJewelsInStones(J, S)
    }
}

func BenchmarkNumJewelsInStonesByArr(b *testing.B) {
    J, S := "aAweio", "aAAbbbbaapawerpaAFWEFAWJOAWRIErwearkaweapwej"
    for i := 0; i < b.N; i++ {
        NumJewelsInStonesByArr(J, S)
    }
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
    s := "fweaoirjaoweryfawo"
    for i := 0; i < b.N; i++ {
        LengthOfLongestSubstring(s)
    }
}

func BenchmarkLengthOfLongestSubstring2(b *testing.B) {
    s := "fweaoirjaoweryfawo"
    for i := 0; i < b.N; i++ {
        LengthOfLongestSubstring2(s)
    }
}

func BenchmarkLengthOfLongestSubstring3(b *testing.B) {
    s := "fweaoirjaoweryfawo"
    for i := 0; i < b.N; i++ {
        LengthOfLongestSubstring3(s)
    }
}

func BenchmarkLengthOfLongestSubstring4(b *testing.B) {
    s := "fweaoirjaoweryfawo"
    for i := 0; i < b.N; i++ {
        LengthOfLongestSubstring4(s)
    }
}

func BenchmarkFourSumCount1(b *testing.B) {
    A := []int{1, 2}
    B := []int{-2, -1}
    C := []int{-1, 2}
    D := []int{0, 2}
    for i := 0; i < b.N; i++ {
        FourSumCount(A, B, C, D)
    }
}

func BenchmarkFourSumCount2(b *testing.B) {
    A := []int{1, 2}
    B := []int{-2, -1}
    C := []int{-1, 2}
    D := []int{0, 2}
    for i := 0; i < b.N; i++ {
        FourSumCount2(A, B, C, D)
    }
}