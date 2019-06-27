package redisutil

import (
	"github.com/go-redis/redis"
	"sync"
)

type ClientManager struct {
	sync.Mutex
	Unused []*redis.Client
}

func (cm *ClientManager) Get() (*redis.Client, bool) {
	var i int
	var c *redis.Client
	cm.Lock()
	for i, c = range cm.Unused {
		if c != nil {
			cm.Unused[i] = nil
			break
		}
	}
	cm.Unlock()

	if c == nil {
		return nil, false
	}

	return c, true
}

func (cm *ClientManager) Register(c *redis.Client) bool {
	cm.Lock()
	defer cm.Unlock()
	for i, uuc := range cm.Unused {
		if uuc == nil {
			cm.Unused[i] = c
			return true
		}
	}
	return false
}
