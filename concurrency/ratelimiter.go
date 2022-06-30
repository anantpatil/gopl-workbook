package main

// Implement a rate limiter - not allow more than 100 requests per second for a client
// implement a isAllowed() func which takes client ID as input and returns true if the request is allowed, false otherwise

import "fmt"
import "sync"

type client struct {
	queue []int
	lck sync.Mutex
}
	
var maxCount = 100
// this map needs to be protected by a lock for cocurrent access
var hashMap = make(map[string]client,0)
func isAllowed(clientId string)  (allowed bool) {
	// take a RLock here to check if client exists
	c , ok := hashMap[clientId]
	if !ok {
		c = client{}
		c.q = append(c.q, time.Now())
		// take a RWLock here and write to the map
		// after taking lock, check again if there is entry for the client
		// if entry not there, create the entry and return True
		// if entry exists, then follow the else {} path below
		hashMap[clientId] = c
		allowed = true
		return
	}

	c.lck.Lock()
	defer c.lck.UnLock()
	now = time.Now()
	if len(c.q) < 100 {
		c.q = append(c.q, now)
		allowed = true
		return
	}
	// purge queue of older requests
	for i := 0 i < len(c.q) && (now - c.q[i]) > 1; i++ {}
	c.q = c.q[i:]
	if len(c.q) < 100 {
		c.q = append(c.q, now)
		allowed = true
	} else {
		fmt.Println("Request not allowed %s", clientId)
	}
}