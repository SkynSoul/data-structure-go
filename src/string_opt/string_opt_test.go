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