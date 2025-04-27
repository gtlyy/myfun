// 功能：实现一个广播器
package myfun

import (
	"fmt"
	"sync"
)

// 广播器结构
type Broadcaster struct {
	mu      sync.Mutex
	clients map[chan string]struct{}
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		clients: make(map[chan string]struct{}),
	}
}

// 订阅新客户端
func (b *Broadcaster) Subscribe() <-chan string {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan string, 10) // 带缓冲的通道
	b.clients[ch] = struct{}{}
	return ch
}

// 向所有客户端广播消息
func (b *Broadcaster) Broadcast(msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for ch := range b.clients {
		select {
		case ch <- msg: // 非阻塞发送
		default:
			fmt.Println("警告：客户端通道已满，丢弃消息")
		}
	}
}

// 关闭所有客户端通道
func (b *Broadcaster) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for ch := range b.clients {
		close(ch)
	}
	b.clients = nil
}
