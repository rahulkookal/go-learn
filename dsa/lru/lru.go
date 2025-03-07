package lru

import "container/list"

type LRU struct {
	cache    map[int]*list.Element
	ll       *list.List
	capacity int
}
type Node struct {
	key   int
	value int
}

func Constructor(capacity int) *LRU {
	return &LRU{
		cache:    make(map[int]*list.Element, capacity),
		ll:       list.New(),
		capacity: capacity,
	}
}

func (lru *LRU) Get(key int) int {
	if elem, found := lru.cache[key]; found {
		lru.ll.MoveToFront(elem)
		return elem.Value.(*Node).value
	}
	return -1
}
func (lru *LRU) Put(key, value int) {
	if elem, found := lru.cache[key]; found {
		elem.Value.(*Node).value = value
		lru.ll.MoveToFront(elem)
		return
	}
	if len(lru.cache) > lru.capacity {
		lru.remove()
	}
	newNode := &Node{
		key:   key,
		value: value,
	}
	llElem := lru.ll.PushFront(newNode)
	lru.cache[key] = llElem
}

func (lru *LRU) remove() {
	rmEl := lru.ll.Back()
	if rmEl != nil {
		lru.ll.Remove(rmEl)
		delete(lru.cache, rmEl.Value.(*Node).key)
	}
}
