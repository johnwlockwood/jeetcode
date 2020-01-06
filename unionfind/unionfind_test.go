package unionfind

import "testing"

func TestUnionFindFind(t *testing.T) {
	n := 10
	items := make([]byte, n)
	for i := 0; i < n; i++ {
		items[i] = byte(i) + 'A'
	}
	u := NewUnionFind(n)
	if got := u.Find(0, 1); got != false {
		t.Errorf("%s and %s should not be connected, but are", string(items[0]), string(items[1]))
	}
	u.Union(0, 7)
	u.Union(7, 1)
	if got := u.Find(7, 1); got != true {
		t.Errorf("%s and %s should be connected, but are not", string(items[7]), string(items[1]))
	}
	if got := u.Find(0, 1); got != true {
		t.Errorf("%s and %s should be connected, but are not", string(items[0]), string(items[1]))
	}
	if got := u.Root(7); got != 0 {
		t.Errorf("Root of %s should be %s, got %s", string(items[7]), string(items[0]), string(items[got]))
	}
	if got := u.Find(8, 1); got != false {
		t.Errorf("%s and %s should not be connected, but are", string(items[8]), string(items[1]))
	}
	u.Union(8, 1)
	if got := u.Root(8); got != 0 {
		t.Errorf("Root of %s should be %s, got %s", string(items[8]), string(items[0]), string(items[got]))
	}
}
