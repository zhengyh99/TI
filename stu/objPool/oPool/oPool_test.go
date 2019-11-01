package oPool

import "testing"

func TestOpool(t *testing.T) {
	pool := GetPool(10)
	for i := 0; i < 12; i++ {
		o, err := pool.GetObject()
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%T\n", o)
		}
	}
}
