package gcache

import "testing"

func TestGCacheGlobalShouldBeAbleToGetValue(t *testing.T) {
	Put("test-get", 1, nil)

	val, err := Get("test-get")
	if err != nil {
		t.Errorf("expect no error, actual %v", err)
	}

	if val != 1 {
		t.Errorf("expect 1, actual %v", val)
	}
}

func TestGCacheGlobalShouldBeAbleToPutValue(t *testing.T) {
	err := Put("test-put", 1, nil)
	if err != nil {
		t.Errorf("expect no error, actual %v", err)
	}
}
