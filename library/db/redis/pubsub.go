package redis

/**
* @DateTime   : 2020/12/30
* @Author     : xumamba
* @Description:
**/

import (
	"sync"

	"github.com/go-redis/redis"
)

type ConnPool struct {
	sync.RWMutex

	ClientPool    map[string]*redis.Client
	PubSubChannel map[string]string

	isClosed bool
}

func (cp *ConnPool) Start() {
	if cp.isClosed != true{
		cp.Close()
	}
	cp.ClientPool = make(map[string]*redis.Client)
	cp.PubSubChannel = make(map[string]string)

}

func (cp *ConnPool) Close() {
	cp.Lock()
	defer cp.Unlock()

	if cp.isClosed == true {
		return
	}
	cp.isClosed = true
	for _, client := range cp.ClientPool {
		_ = client.Close()
	}
}
