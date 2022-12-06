package main

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
	head = &Node{key: -2, value: -2}
	tail = &Node{key: -1, value: -1}
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
		if currentCapacity >= l.capacity {
			fmt.Println("capacity exceeded. removing old cache value")
			temp := tail.prev
			temp.prev.next = temp.next
			temp.next.prev = temp.prev
			temp.next = nil
			temp.prev = nil

			currentCapacity--
			delete(cache, temp.value)
		}

		newNode := &Node{
			key:   newValue,
			value: newValue,
		}

		head.next.prev = newNode
		newNode.next = head.next
		head.next = newNode
		newNode.prev = head

		cache[newValue] = newNode
		currentCapacity++
	}
}

func (l *LRUCache) Get(value int) (int, error) {
	entry, ok := cache[value]

	if !ok {
		return 0, fmt.Errorf("value not available in cache")
	}

	// move node to the beginning
	entry.prev.next = entry.next
	entry.next.prev = entry.prev

	entry.next = head.next
	head.next.prev = entry
	head.next = entry
	entry.prev = head

	return entry.key, nil
}

func (l *LRUCache) Print() {
	curr := head
	for curr != nil {
		fmt.Printf("%v -> ", curr.key)
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	lrucache := NewLRUCache(5)
	lrucache.Add(1)
	lrucache.Add(2)
	lrucache.Add(3)
	lrucache.Add(4)
	lrucache.Add(5)
	lrucache.Get(2)
	lrucache.Print()
}
