package linkedlist

// LeetCode problem 146. LRU Cache
// https://leetcode.com/problems/lru-cache/
// Medium
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
		if node == this.usageList.head {
			return entry.value
		}

		// promote the most recently used key's node to head
		this.promoteMostRecentlyUsedNodeToHead(node)
		return entry.value
	}
	return -1
}

func (this *LRUCache) promoteMostRecentlyUsedNodeToHead(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if this.usageList.tail == node {
		this.usageList.tail = this.usageList.tail.prev
	}
	node.next = this.usageList.head
	if this.usageList.head != nil {
		this.usageList.head.prev = node
	}
	node.prev = nil
	this.usageList.head = node
}

func (this *LRUCache) Put(key int, value int) {
	if entry, exists := this.cache[key]; exists {
		entry.value = value
		return
	}

	this.usageList.size++
	if this.usageList.size > this.usageList.cap {
		this.ejectLeastUsedEntry()
	}

	newNode := &Node{key: key}
	this.promoteMostRecentlyUsedNodeToHead(newNode)
	//newNode := this.updateHead(key, value)
	this.cache[key] = &Entry{value, newNode}

}

func (this *LRUCache) updateHead(key int, value int) *Node {
	newNode := &Node{key: key, prev: nil, next: nil}
	newNode.prev = nil
	newNode.next = this.usageList.head
	if this.usageList.head != nil {
		this.usageList.head.prev = newNode
	}
	this.usageList.head = newNode
	if this.usageList.tail == nil {
		this.usageList.tail = this.usageList.head
	}
	return newNode
}

func (this *LRUCache) ejectLeastUsedEntry() {
	luKey := this.usageList.tail.key
	toEject := this.usageList.tail
	this.usageList.tail = toEject.prev
	if this.usageList.tail != nil {
		this.usageList.tail.next = nil
	}
	toEject.prev = nil
	this.usageList.size--
	delete(this.cache, luKey)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
