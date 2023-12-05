package gcache

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrKeyAlreadyExists = errors.New("error key already exists")
	ErrKeyNotFound      = errors.New("error key not found")
	ErrRenewFailed      = errors.New("error renew failed")
	ErrKeyExpired       = errors.New("error key expired")
)

var (
	Day            = 24 * time.Hour
	Year           = 365 * Day
	InfiniteExpiry = time.Now().AddDate(9999, 0, 0)
	InfiniteExtend = 9999 * Year
)

type Cache struct {
	items map[string]CacheItem
}

func New() Cache {
	return Cache{
		items: map[string]CacheItem{},
	}
}

func (c *Cache) Put(key string, val interface{}, options *ItemOptions) error {
	_, found := c.items[key]
	if found {
		return ErrKeyAlreadyExists
	}

	c.items[key] = newCacheItem(val, options)

	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	item, found := c.items[key]
	if !found {
		return nil, ErrKeyNotFound
	}

	var err error = nil

	if item.expiry.Before(time.Now()) {
		item.lock.Lock()
		err = item.renew()
		item.expiry = item.expiry.Add(item.extendDuration)
		item.lock.Unlock()
	}

	item.lock.RLock()
	defer item.lock.RUnlock()
	return item.value, err
}

type CacheItem struct {
	value          interface{}
	extendDuration time.Duration
	expiry         time.Time
	valueReNewer   ReNewerFunc
	lock           sync.RWMutex
}

type ReNewerFunc func() (interface{}, error)

func newCacheItem(val interface{}, options *ItemOptions) CacheItem {
	if options == nil {
		options = &ItemOptions{
			Expiry:         InfiniteExpiry,
			ExtendDuration: InfiniteExtend,
		}
	}

	return CacheItem{
		value:          val,
		expiry:         options.Expiry,
		extendDuration: options.ExtendDuration,
		valueReNewer:   options.RenewFunc,
	}
}

func (ci *CacheItem) renew() error {
	if ci.valueReNewer == nil {
		return ErrKeyExpired
	}

	newVal, err := ci.valueReNewer()
	if err != nil {
		return ErrRenewFailed
	}

	ci.value = newVal
	return nil
}

type ItemOptions struct {
	Expiry         time.Time
	ExtendDuration time.Duration
	RenewFunc      ReNewerFunc
}
