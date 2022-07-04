package main

// Implement a rate limiter - not allow more than 100 requests per second for a client
// implement a isAllowed() func which takes client ID as input and returns true if the request is allowed, false otherwise

import (
	"fmt"
	"sync"
	"time"
)

type client struct {
	sync.Mutex
	q []time.Time
}

const RATE_LIMIT = 100

var clients = make(map[string]*client, 0)

// lock for clients
var lock sync.RWMutex

func IsRequestAllowed(clientId string) (allowed bool) {
	// take a RLock here to check if client exists
	lock.RLock()
	c, ok := clients[clientId]
	lock.RUnlock()
	if !ok {
		// take a RWLock here and write to the map
		// after taking lock, check again if there is entry for the client
		// if entry not there, create the entry and return True
		// if entry exists, then follow the else {} path below
		lock.Lock()
		c, ok = clients[clientId]
		if !ok {
			c = &client{}
			c.q = append(c.q, time.Now())
			clients[clientId] = c
			allowed = true
			lock.Unlock()
			return
		}
		lock.Unlock()
	}

	// handle concurrent requests per client
	c.Lock()
	defer c.Unlock()
	if len(c.q) < RATE_LIMIT {
		c.q = append(c.q, time.Now())
		allowed = true
		return
	}

	// purge queue of older requests
	idx := 0
	now := time.Now()
	for idx < len(c.q) && now.Sub(c.q[idx]).Milliseconds() > 1000 {
		idx++
	}
	c.q = c.q[idx:]
	if len(c.q) < RATE_LIMIT {
		c.q = append(c.q, now)
		allowed = true
	} else {
		fmt.Printf("Request not allowed %s\n", clientId)
		fmt.Printf("queue len: %d\n", len(c.q))
	}
	return
}

func main() {
	for i := 0; i < 102; i++ {
		go func(i int) {
			fmt.Printf("Allowed(%d): [%v] %v\n", i, time.Now(), IsRequestAllowed("123456789"))
			time.Sleep(1 * time.Second)
			fmt.Printf("Allowed2(%d): [%v] %v\n", i, time.Now(), IsRequestAllowed("123456789"))
		}(i)
	}
	quit := make(chan struct{})
	<-quit
}
