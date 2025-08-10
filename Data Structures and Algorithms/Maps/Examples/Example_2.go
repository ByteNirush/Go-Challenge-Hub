// Example 2: Implementing a Simple Cache
// This example demonstrates how to use a map to implement a simple cache with a fixed size.

package main

import (
    "fmt"
    "time"
)

// SimpleCache represents a simple cache implementation.
type SimpleCache struct {
    data      map[string]interface{}
    capacity  int
    evictionPolicy string // "LRU" (Least Recently Used), "FIFO" (First-In-First-Out)
    accessTimes map[string]time.Time // For LRU eviction
    queue []string // For FIFO eviction
}

// NewSimpleCache creates a new SimpleCache instance.
func NewSimpleCache(capacity int, evictionPolicy string) *SimpleCache {
    return &SimpleCache{
        data:      make(map[string]interface{}),
        capacity:  capacity,
        evictionPolicy: evictionPolicy,
        accessTimes: make(map[string]time.Time),
        queue: make([]string, 0, capacity),
    }
}

// Get retrieves a value from the cache.
func (c *SimpleCache) Get(key string) (interface{}, bool) {
    value, ok := c.data[key]
    if ok {
        if c.evictionPolicy == "LRU"{
            c.accessTimes[key] = time.Now() // Update access time
        }
        return value, true
    }
    return nil, false
}

// Put adds a value to the cache.
func (c *SimpleCache) Put(key string, value interface{}) {
    if _, ok := c.data[key]; !ok && len(c.data) >= c.capacity {
        c.evict()
    }
    c.data[key] = value
    if c.evictionPolicy == "LRU"{
        c.accessTimes[key] = time.Now()
    } else if c.evictionPolicy == "FIFO"{
        if _, ok := c.data[key]; !ok{
            c.queue = append(c.queue, key)
        }
    }
}

// evict removes an element from the cache based on the eviction policy.
func (c *SimpleCache) evict() {
    if c.evictionPolicy == "LRU"{
        var lruKey string
        var lruTime time.Time
        first := true

        for key, accessTime := range c.accessTimes {
            if first || accessTime.Before(lruTime) {
                lruKey = key
                lruTime = accessTime
                first = false
            }
        }

        delete(c.data, lruKey)
        delete(c.accessTimes, lruKey)
    } else if c.evictionPolicy == "FIFO"{
        if len(c.queue) > 0{
            oldestKey := c.queue[0]
            c.queue = c.queue[1:]
            delete(c.data, oldestKey)
        }
    }
}

func main() {
    cache := NewSimpleCache(3, "LRU")

    cache.Put("key1", "value1")
    cache.Put("key2", "value2")
    cache.Put("key3", "value3")

    value, ok := cache.Get("key1")
    if ok {
        fmt.Println("Key1:", value)
    }

    cache.Put("key4", "value4") // This will evict the least recently used key

    value, ok = cache.Get("key2")
    if ok {
        fmt.Println("Key2:", value)
    } else {
        fmt.Println("Key2: Not found")
    }

    value, ok = cache.Get("key1")
    if ok {
        fmt.Println("Key1:", value)
    } else {
        fmt.Println("Key1: Not found")
    }
}
// This example demonstrates a simple cache implementation using a map. The SimpleCache struct contains a map to store the cached data, a capacity to limit the cache size, and an eviction policy to determine which element to remove when the cache is full. The LRU eviction policy removes the least recently used element. The FIFO eviction policy removes the first element that was added.

