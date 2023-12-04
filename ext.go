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

func (c *Cache) PutInt64(key string, val int64, options *ItemOptions) error {
	return c.Put(key, val, options)
}

func (c *Cache) GetInt64(key string) (int64, error) {
	val, err := c.Get(key)
	if err != nil {
		return 0, err
	}

	return val.(int64), nil
}

func (c *Cache) PutFloat64(key string, val float64, options *ItemOptions) error {
	return c.Put(key, val, options)
}

func (c *Cache) GetFloat64(key string) (float64, error) {
	val, err := c.Get(key)
	if err != nil {
		return 0, err
	}

	return val.(float64), nil
}

func (c *Cache) PutBool(key string, val bool, options *ItemOptions) error {
	return c.Put(key, val, options)
}

func (c *Cache) GetBool(key string) (bool, error) {
	val, err := c.Get(key)
	if err != nil {
		return false, err
	}

	return val.(bool), nil
}

func (c *Cache) PutBytes(key string, val []byte, options *ItemOptions) error {
	return c.Put(key, val, options)
}

func (c *Cache) GetBytes(key string) ([]byte, error) {
	val, err := c.Get(key)
	if err != nil {
		return nil, err
	}

	return val.([]byte), nil
}
