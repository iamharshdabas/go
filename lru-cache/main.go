package main

func main() {
}

type Node[K comparable, V any] struct {
	left  *Node[K, V]
	key   K
	value V
	right *Node[K, V]
}

type LRUCache[K comparable, V any] struct {
	head      *Node[K, V]
	tail      *Node[K, V]
	length    int
	maxLength int
	hashMap   map[K]*Node[K, V]
}

func newLRUCache[K comparable, V any](maxLength int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		head:      nil,
		tail:      nil,
		length:    0,
		maxLength: maxLength,
		hashMap:   make(map[K]*Node[K, V], maxLength),
	}
}

func (l *LRUCache[K, V]) get(key K) (V, bool) {
	var zeroValue V

	node := l.search(key)
	if node == nil {
		return zeroValue, false
	}

	l.put(node)
	return node.value, true
}

func (l *LRUCache[K, V]) add(key K, value V) {
	l.put(&Node[K, V]{
		key:   key,
		value: value,
	})
}

func (l *LRUCache[K, V]) put(node *Node[K, V]) {
	if l.length == 0 {
		l.head = node
		l.tail = node
		l.length++
		l.hashMap[node.key] = node
		return
	}

	nodeInList := l.search(node.key)
	if nodeInList != nil {
		l.delete(nodeInList)
	}

	if l.length == l.maxLength {
		l.delete(l.tail)
	}

	if l.length == 0 {
		l.head = node
		l.tail = node
		l.length++
		l.hashMap[node.key] = node
		return
	}

	node.right = l.head
	l.head.left = node
	l.head = node

	l.length++
	l.hashMap[node.key] = node
}

func (l *LRUCache[K, V]) delete(node *Node[K, V]) {
	var (
		prevNode *Node[K, V]
		nextNode *Node[K, V]
	)

	if node == nil {
		return
	}

	if node.left != nil {
		prevNode = node.left
		prevNode.right = nil
		node.left = nil
	}

	if node.right != nil {
		nextNode = node.right
		nextNode.left = nil
		node.right = nil
	}

	if prevNode != nil && nextNode != nil {
		prevNode.right = nextNode
		nextNode.left = prevNode
	}

	if node == l.head {
		if nextNode != nil {
			l.head = nextNode
		} else {
			l.head = nil
		}
	}

	if node == l.tail {
		if prevNode != nil {
			l.tail = prevNode
		} else {
			l.tail = nil
		}
	}

	l.length--
	delete(l.hashMap, node.key)
}

func (l *LRUCache[K, V]) search(key K) *Node[K, V] {
	return l.hashMap[key]
}
