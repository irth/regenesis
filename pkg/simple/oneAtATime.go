package simple

import "sync"

type OneAtATime struct {
	reserved bool
	lock     sync.Mutex
}

func (o OneAtATime) Lock() bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	if !o.reserved {
		o.reserved = true
		return true
	}
	return false
}

func (o OneAtATime) Unlock() {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.reserved = false
}
