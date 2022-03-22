package lfu_cache

type LFUCache struct {
	freqList  *freqList
	freqMap   map[int]*freqNode // freq - node
	cacheList *cacheList
	cacheMap  map[int]*cacheNode // key - node
	capacity  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		freqList:  newFreqList(),
		freqMap:   make(map[int]*freqNode),
		cacheList: newCacheList(),
		cacheMap:  make(map[int]*cacheNode),
		capacity:  capacity,
	}
}

func (lfu *LFUCache) Get(key int) int {
	if lfu.capacity == 0 {
		return -1
	}

	node, ok := lfu.cacheMap[key]
	if !ok { // key 不存在
		return -1
	}

	// key 存在
	node.remove()

	freq := node.frequency
	fNode := lfu.freqMap[freq]

	newFreqNode, ok := lfu.freqMap[freq+1]
	if !ok { // 不存在 freq = n+1 的节点
		newFreqNode = &freqNode{
			frequency: freq + 1,
			data:      newCacheList(),
		}

		fNode.addBehind(newFreqNode)
		lfu.freqMap[freq+1] = newFreqNode
	}

	node.frequency++
	newFreqNode.data.addToHead(node)

	if fNode.data.isEmpty() { // 频率节点(freq = n)为空，后删除是为了帮助freq = n+1的频率节点在新增时定位
		fNode.remove()
		delete(lfu.freqMap, freq)
	}

	return node.value
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}

	node, ok := lfu.cacheMap[key]
	if ok {
		lfu.Get(key)
		node.value = value

		return
	}

	if len(lfu.cacheMap) >= lfu.capacity {
		fNode := lfu.freqList.head.next

		delNode := fNode.data.tail.prev
		delNode.remove()
		delete(lfu.cacheMap, delNode.key)

		if fNode.data.isEmpty() && fNode.frequency > 1 {
			delete(lfu.freqMap, fNode.frequency)
		}
	}

	fNode, ok := lfu.freqMap[1]
	if !ok {
		fNode = &freqNode{
			frequency: 1,
			data:      newCacheList(),
		}

		lfu.freqList.addToHead(fNode)
		lfu.freqMap[1] = fNode
	}

	newCacheNode := &cacheNode{
		key:       key,
		value:     value,
		frequency: 1,
	}

	fNode.data.addToHead(newCacheNode)
	lfu.cacheMap[key] = newCacheNode
}

type freqList struct {
	head *freqNode
	tail *freqNode
}

type freqNode struct {
	frequency int
	data      *cacheList
	prev      *freqNode
	next      *freqNode
}

func newFreqList() *freqList {
	headNode := &freqNode{}
	tailNode := &freqNode{}

	headNode.next = tailNode
	tailNode.prev = headNode

	return &freqList{
		head: headNode,
		tail: tailNode,
	}
}

func (f *freqNode) remove() {
	f.prev.next = f.next
	f.next.prev = f.prev
}

func (f *freqNode) addBehind(node *freqNode) {
	next := f.next

	f.next = node
	node.next = next

	next.prev = node
	node.prev = f
}

func (fl *freqList) addToHead(node *freqNode) {
	next := fl.head.next

	fl.head.next = node
	node.next = next

	next.prev = node
	node.prev = fl.head
}

type cacheList struct {
	head *cacheNode
	tail *cacheNode
}

type cacheNode struct {
	key       int
	value     int
	frequency int
	prev      *cacheNode
	next      *cacheNode
}

func newCacheList() *cacheList {
	headNode := &cacheNode{}
	tailNode := &cacheNode{}

	headNode.next = tailNode
	tailNode.prev = headNode

	return &cacheList{
		head: headNode,
		tail: tailNode,
	}
}

func (c *cacheNode) remove() {
	c.prev.next = c.next
	c.next.prev = c.prev
}

func (cl *cacheList) addToHead(node *cacheNode) {
	next := cl.head.next

	cl.head.next = node
	node.next = next

	next.prev = node
	node.prev = cl.head
}

func (cl *cacheList) isEmpty() bool {
	return cl.head.next == cl.tail
}
