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