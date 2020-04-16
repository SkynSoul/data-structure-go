package double_pointer

import "testing"

func TestReverseString(t *testing.T) {
	arr := []byte{'h', 'e', 'l', 'l', 'o'}
	t.Logf("src str arr: %q\n", arr)
	ReverseString(arr)
	t.Logf("after reverse: %q\n", arr)
}