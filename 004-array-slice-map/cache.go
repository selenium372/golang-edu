package main

import "sync"

type MyCache struct {
	sharedCache map[string]string
	cacheMutex  sync.RWMutex
}

type Cache interface {
	Get(key string) string
	Set(key string, value string)
}

func (c *MyCache) Get(key string) string {
	var s string
	c.cacheMutex.RLock()
	s = c.sharedCache[key]
	c.cacheMutex.RUnlock()
	return s
}

func (c *MyCache) Set(key string, value string) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()
	c.sharedCache[key] = value
}
