package linkedlist

import (
	"fmt"
	"strconv"
	"strings"
)

// LeetCode problem 146. LRU Cache
// https://leetcode.com/problems/lru-cache/
// Medium

// Using double linked list (not circular) and hash map
type (
	LRUCache struct {
		cache     map[int]*Entry
		usageList UsageList
	}

	Entry struct {
		value int
		node  *Node
	}

	UsageList struct {
		head,
		tail *Node
		size,
		cap int
	}

	Node struct {
		key int
		prev,
		next *Node
	}
)

func NewLRUCache(capacity int) LRUCache {
	cache := LRUCache{
		cache: make(map[int]*Entry, capacity),
		usageList: UsageList{
			cap: capacity,
		},
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	if entry, exists := this.cache[key]; exists {
		node := entry.node
		if node != this.usageList.head {
			// promote the most recently used key's node to head
			this.promoteMostRecentlyUsedNodeToHead(node)
		}
		fmt.Printf("GET(%d), Return:%d, Cache:%#v\n", key, entry.value, printCache(this.cache))

		return entry.value
	}
	fmt.Printf("GET(%d), Return:%d, Cache:%#v\n", key, -1, printCache(this.cache))
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	var entry *Entry
	var exists bool
	var node *Node
	if entry, exists = this.cache[key]; exists {
		node = entry.node
	} else {
		if this.usageList.size+1 > this.usageList.cap {
			this.ejectLeastUsedEntry()
		}
		this.usageList.size++
		node = &Node{key: key}
	}

	if node != this.usageList.head {
		this.promoteMostRecentlyUsedNodeToHead(node)
	}
	this.cache[key] = &Entry{value, node}
	fmt.Printf("PUT(%d,%d),Cache:%#v\n", key, value, printCache(this.cache))
}

func printCache(cache map[int]*Entry) string {
	var sb strings.Builder
	for k, v := range cache {
		sb.WriteString(strconv.Itoa(k) + "," + strconv.Itoa(v.value))
	}
	return "{" + sb.String() + "}"
}

func (this *LRUCache) promoteMostRecentlyUsedNodeToHead(node *Node) {
	// if prev exists, connect prev with next
	if node.prev != nil {
		node.prev.next = node.next
	}
	// 	if next exists, connect next with previous
	if node.next != nil {
		node.next.prev = node.prev
	}
	// if node is the tail, new tail is tail.prev
	if this.usageList.tail == node {
		this.usageList.tail = this.usageList.tail.prev
	}
	// put node at head
	node.next = this.usageList.head
	if this.usageList.head != nil {
		this.usageList.head.prev = node
	}
	node.prev = nil
	this.usageList.head = node
	// if there's no next, it means the node is the tail
	if this.usageList.head.next == nil {
		this.usageList.tail = this.usageList.head
	}
}

func (this *LRUCache) ejectLeastUsedEntry() {
	delete(this.cache, this.usageList.tail.key)
	this.usageList.size--
	if this.usageList.size <= 0 {
		this.usageList.head = nil
		this.usageList.tail = nil
	} else {
		this.usageList.tail.prev.next = nil
		this.usageList.tail = this.usageList.tail.prev
	}
	/*
		luKey := this.usageList.tail.key
		   toEject := this.usageList.tail
		   // The new tail is the previous of the one to eject
		   this.usageList.tail = toEject.prev

		   	if this.usageList.tail != nil {
		   		this.usageList.tail.next = nil
		   	}

		   toEject.prev = nil

		   	if this.usageList.size <= 0 {
		   		this.usageList.head = nil
		   	}

		   delete(this.cache, luKey)
	*/
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
