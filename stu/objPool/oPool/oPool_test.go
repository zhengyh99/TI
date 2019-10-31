package oPool

import "testing"

func TestOpool(t *testing.T) {
	pool := GetPool(10)
	for i := 0; i < 10; i++ {
		if o, err := pool.GetObject(); err != nil {
			t.Error(err)
		} else {
			t.Logf("%T\n", o)
		}
	}
}
