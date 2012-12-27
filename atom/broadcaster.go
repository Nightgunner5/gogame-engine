package atom

import (
	"sync"
)

type Broadcaster interface {
	Broadcast(Signal) (count int)
	Subscribe(Kind, chan<- Signal)
	Unsubscribe(Kind, chan<- Signal)
}

type broadcaster struct {
	subscribers map[Kind]map[chan<- Signal]bool
	sync.RWMutex
}

func NewBroadcaster() Broadcaster {
	return &broadcaster{
		subscribers: make(map[Kind]map[chan<- Signal]bool),
	}
}

func (b *broadcaster) Broadcast(s Signal) (count int) {
	b.RLock()
	for sub := range b.subscribers[s.Kind()] {
		select {
		case sub <- s:
			count++
		default:
		}
	}
	b.RUnlock()
	return
}

func (b *broadcaster) Subscribe(k Kind, c chan<- Signal) {
	b.Lock()
	if b.subscribers[k] == nil {
		b.subscribers[k] = make(map[chan<- Signal]bool)
	}
	b.subscribers[k][c] = true
	b.Unlock()
}

func (b *broadcaster) Unsubscribe(k Kind, c chan<- Signal) {
	b.Lock()
	if b.subscribers[k] == nil {
		b.Unlock()
		return
	}
	delete(b.subscribers[k], c)
	if len(b.subscribers[k]) == 0 {
		delete(b.subscribers, k)
	}
	b.Unlock()
}
