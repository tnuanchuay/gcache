package gcache

func (c *Cache) PutString(key string, val string, options *ItemOptions) error {
	return c.Put(key, val, options)
}

func (c *Cache) GetString(key string) (string, error) {
	val, err := c.Get(key)
	if err != nil {
		return "", err
	}

	return val.(string), nil
}

func (c *Cache) PutInt(key string, val int, options *ItemOptions) error {
	return c.Put(key, val, options)
}

func (c *Cache) GetInt(key string) (int, error) {
	val, err := c.Get(key)
	if err != nil {
		return 0, err
	}

	return val.(int), nil
}
