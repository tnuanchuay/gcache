package gcache

import (
	"errors"
	"testing"
	"time"
)

var (
	TestExpectActualFormat = "expect %v, actual %v"
)

func TestGCacheShouldReturnErrRenewFailedWhenRenewFuncReturnError(t *testing.T) {
	cache := New()
	cache.items["test-key"] = CacheItem{
		valueReNewer: func() (interface{}, error) {
			return nil, errors.New("error")
		},
	}

	_, err := cache.Get("test-key")
	if !errors.Is(err, ErrRenewFailed) {
		t.Errorf(TestExpectActualFormat, ErrRenewFailed, err)
	}
}

func TestGCacheShouldReturnErrExpiredWhenGetExpiredKey(t *testing.T) {
	cache := New()
	cache.items["test-key"] = CacheItem{
		value: 1,
	}

	_, err := cache.Get("test-key")
	if !errors.Is(err, ErrKeyExpired) {
		t.Errorf(TestExpectActualFormat, ErrKeyExpired, err)
	}
}

func TestGCacheShouldReturnKeyNotFound(t *testing.T) {
	cache := New()

	_, err := cache.Get("test-key")
	if !errors.Is(err, ErrKeyNotFound) {
		t.Errorf(TestExpectActualFormat, ErrKeyNotFound, err)
	}
}

func TestGCacheShouldReturnErrKeyAlreadyExistsWhenPutExistingKey(t *testing.T) {
	cache := New()
	cache.items["test-key"] = CacheItem{}

	err := cache.Put("test-key", 1, nil)
	if !errors.Is(err, ErrKeyAlreadyExists) {
		t.Errorf(TestExpectActualFormat, ErrKeyAlreadyExists, err)
	}
}

func TestGCacheShouldBeAbleToReNewValueAndUpdateExpiry(t *testing.T) {
	cache := New()
	cache.Put("test-key", 1, &ItemOptions{
		Expiry:         time.Now().Add(1 * time.Second * -1),
		ExtendDuration: 5 * time.Minute,
		RenewFunc: func() (interface{}, error) {
			return 2, nil
		},
	})

	v, err := cache.Get("test-key")
	if err != nil {
		t.Error(err)
	}

	if v != 2 {
		t.Errorf(TestExpectActualFormat, 2, v)
	}
	timeDiff := time.Now().Sub(cache.items["test-key"].expiry)
	if timeDiff > 5*time.Minute {
		t.Errorf(TestExpectActualFormat, 5*time.Minute, timeDiff)
	}
}

func TestGCacheShouldBeAbleToGetExistingKey(t *testing.T) {
	cache := New()
	cache.items["test-key"] = CacheItem{
		value:  1,
		expiry: InfiniteExpiry,
	}

	val, err := cache.Get("test-key")
	if err != nil {
		t.Error(err)
	}

	if val != 1 {
		t.Errorf(TestExpectActualFormat, 1, val)
	}
}

func TestGCacheShouldBeAbleToPutNewKey(t *testing.T) {
	cache := New()
	err := cache.Put("test-key", 1, nil)
	if err != nil {
		t.Error(err)
	}

	item := cache.items["test-key"]
	if item.value != 1 {
		t.Errorf(TestExpectActualFormat, 1, item.value)
	}
}

func TestGCacheShouldReturnNewInstance(t *testing.T) {
	cache := New()
	if len(cache.items) != 0 {
		t.Errorf(TestExpectActualFormat, 0, len(cache.items))
	}
}
