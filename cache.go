package cache

import "time"

var (
	zeroTime = *new(time.Time)
)

type Cache struct {
	cache   map[string]string
	expired map[string]time.Time
}

func NewCache() Cache {
	return Cache{make(map[string]string), make(map[string]time.Time)}
}

func (c Cache) Get(key string) (string, bool) {
	a, b := c.cache[key]
	expired := c.expired[key]
	if !b || (!expired.IsZero() && expired.Before(time.Now())) {
		return "", false
	}

	return a, true
}

func (c Cache) Put(key, value string) {
	c.PutTill(key, value, zeroTime)
}

func (c Cache) Keys() []string {
	now := time.Now()
	var result []string
	for k, v := range c.expired {
		if v.IsZero() || now.Before(v) {
			result = append(result, k)
		}
	}
	return result
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.cache[key] = value
	c.expired[key] = deadline
}
