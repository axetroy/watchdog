package socket

import (
	"sync"
)

// SafeCounter is safe to use concurrently.
type SocketPool struct {
	mu sync.Mutex
	v  map[string]*Socket
}

func (p *SocketPool) Remove(uid string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.v, uid)
}

func (p *SocketPool) Add(uid string, conn *Socket) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.v[uid] = conn
}

func (p *SocketPool) Broadcast(v Data) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, conn := range p.v {
		_ = conn.WriteJSON(v)
	}
}

func (p *SocketPool) BroadcastTo(uuid string, v Data) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, conn := range p.v {
		if uuid == conn.UUID {
			_ = conn.WriteJSON(v)
		}
	}
}

var Pool = &SocketPool{
	v: make(map[string]*Socket),
}
