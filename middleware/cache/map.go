package cache

import (
	"sync"
	"sync/atomic"
	"time"
)

type IPInfo struct {
	expire time.Time
	cnt    atomic.Value
}

type cache struct {
	ips sync.Map
	// janitor   *janitor
}

// func (c *cache) Set(ip string, expire time.Duration) {
//
// 	c.ips.Store(ip, IPInfo{
// 		expire: time.Now().Add(expire),
// 		cnt:    atomic.Value{},
// 	})
// }
//
// func (c *cache) Get(k string) (interface{}, bool) {
//
// }
