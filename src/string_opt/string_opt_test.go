package string_opt

import (
    "reflect"
    "testing"
)

func TestOptString(t *testing.T) {
    str := "123456"
    t.Logf("str type is %s\n", reflect.TypeOf(str).Name())
    // 不能直接操作字符串
    //str[0] = '0'
    //t.Logf("str is %s\n", str)
}

func TestAddBinary(t *testing.T) {
    a, b := "11", "1"
    t.Logf("%s + %s = %s\n", a, b, AddBinary(a, b))
    a, b = "1010", "1011"
    t.Logf("%s + %s = %s\n", a, b, AddBinary(a, b))
    a, b = "1010001", "1011"
    t.Logf("%s + %s = %s\n", a, b, AddBinary(a, b))
}

func TestStrStr(t *testing.T) {
    haystack, needle := "hello", "ll"
    t.Logf("haystack is %s, needle is %s, the first index is %d\n", haystack, needle, StrStr(haystack, needle))
    haystack, needle = "aaaaa", "bba"
    t.Logf("haystack is %s, needle is %s, the first index is %d\n", haystack, needle, StrStr(haystack, needle))
    haystack, needle = "aaaaa", ""
    t.Logf("haystack is %s, needle is %s, the first index is %d\n", haystack, needle, StrStr(haystack, needle))
    haystack, needle = "aaaaa", "a"
    t.Logf("haystack is %s, needle is %s, the first index is %d\n", haystack, needle, StrStr(haystack, needle))
    haystack, needle = "", ""
    t.Logf("haystack is %s, needle is %s, the first index is %d\n", haystack, needle, StrStr(haystack, needle))
}

func TestLongestCommonPrefix(t *testing.T) {
    strs := []string{"flower", "flow", "flight"}
    t.Logf("strs is %v, the longest common prefix is %s\n", strs, LongestCommonPrefix(strs))
    strs = []string{"dog", "racecar", "car"}
    t.Logf("strs is %v, the longest common prefix is %s\n", strs, LongestCommonPrefix(strs))
    strs = []string{"car"}
    t.Logf("strs is %v, the longest common prefix is %s\n", strs, LongestCommonPrefix(strs))
    strs = []string{"c", "c"}
    t.Logf("strs is %v, the longest common prefix is %s\n", strs, LongestCommonPrefix(strs))
}