package cache

import "sync"

type Entry struct {
	Value   any
	Version uint64
}
type Cache struct {
	mu      sync.RWMutex
	data    map[string]Entry
	version uint64
}

func New() *Cache { return &Cache{data: map[string]Entry{}} }
func (c *Cache) Put(k string, v any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.version++
	c.data[k] = Entry{v, c.version}
}
func (c *Cache) Get(k string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[k]
	return v.Value, ok
}
func (c *Cache) Invalidate(k string)      { c.mu.Lock(); defer c.mu.Unlock(); delete(c.data, k) }
func (c *Cache) InvalidateTeam(id string) { c.Invalidate("launch:" + id) }
func (c *Cache) Clear()                   { c.mu.Lock(); defer c.mu.Unlock(); c.data = map[string]Entry{} }
func (c *Cache) Size() int                { c.mu.RLock(); defer c.mu.RUnlock(); return len(c.data) }
