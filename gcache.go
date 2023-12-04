package gcache

var (
	cache Cache = New()
)

func Put(key string, val interface{}, options *ItemOptions) error {
	return cache.Put(key, val, options)
}

func Get(key string) (interface{}, error) {
	return cache.Get(key)
}
