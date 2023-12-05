package gcache

import "testing"

func TestShouldBeAbleToPutString(t *testing.T) {
	cache := New()
	err := cache.PutString("test-string", "test", nil)
	if err != nil {
		t.Errorf("expect no error, actual %v", err)
	}

	val, err := cache.GetString("test-string")
	if err != nil {
		t.Errorf("expect no error, actual %v", err)
	}

	if val != "test" {
		t.Errorf("expect test, actual %v", val)
	}
}
