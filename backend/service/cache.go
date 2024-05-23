package service

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Item struct {
	Key        string      `json:"key"`
	Value      interface{} `json:"value"`
	Expiration int64       `json:"expiration"` // Expiration time in seconds
}

type CacheService struct {
	items   map[string]Item
	mu      sync.RWMutex
	Clients map[*websocket.Conn]bool
}

func NewCacheService() *CacheService {
	cache := &CacheService{
		items: make(map[string]Item),
	}
	go cache.startEvictionTicker()
	return cache
}

func (c *CacheService) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[key]
	if !found || time.Now().UnixNano() > item.Expiration {
		return nil, false
	}
	return item.Value, true
}

func (c *CacheService) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expiration := time.Now().Add(duration).UnixNano()
	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
	}
}

func (c *CacheService) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *CacheService) startEvictionTicker() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		c.evictExpiredItems()
	}
}

func (c *CacheService) evictExpiredItems() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now().UnixNano()
	for k, v := range c.items {
		if now > v.Expiration {
			delete(c.items, k)
		}
	}
}

func (c *CacheService) RegisterClient(client *websocket.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Clients[client] = true
}

func (c *CacheService) UnregisterClient(client *websocket.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Clients, client)
}

func (c *CacheService) GetItems() map[string]Item {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.items
}
