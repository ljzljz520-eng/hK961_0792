package cache

import "testing"

func TestCacheInvalidation(t *testing.T) {
	c := New()
	c.Put("x", 1)
	if c.Size() != 1 {
		t.Fatal()
	}
	c.Invalidate("x")
	if _, ok := c.Get("x"); ok {
		t.Fatal()
	}
}
