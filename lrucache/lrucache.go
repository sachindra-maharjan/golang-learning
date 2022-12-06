package lrucache

import (
	"fmt"
)

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
}

var cache = make(map[int]*Node)
var currentCapacity int
var head *Node
var tail *Node

func NewLRUCache(capacity int) *LRUCache {
	head = &Node{}
	tail = &Node{}
	head.next = tail
	tail.prev = head
	return &LRUCache{capacity: capacity}
}

func (l *LRUCache) Add(newValue int) {
	entry, ok := cache[newValue]
	if ok {
		entry.next = head.next
		entry.prev = head.prev
		head.next = entry
	} else {
		// remove tail when maximum capacity exceeds
		if currentCapacity > l.capacity {
			temp := tail.prev
			temp.next = nil
			tail.prev = nil
			tail = temp
			currentCapacity--
		}

		newNode := &Node{
			key:   newValue,
			value: newValue,
			prev:  head,
			next:  head.next,
		}

		head.next.prev = newNode
		head.next = newNode

		currentCapacity++
	}
}

func (l *LRUCache) Get(value int) (int, error) {
	entry, ok := cache[value]

	if !ok {
		return 0, fmt.Errorf("value not available in cache")
	}

	// move node to the beginning
	entry.next = head.next
	entry.prev = head.prev
	head.next = entry

	return entry.key, nil
}

func main() {

}
