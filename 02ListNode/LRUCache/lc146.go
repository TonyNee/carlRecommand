package main

import (
	"carlRecommend/02ListNode/LRUCache/lru"
)

func main() {
	obj := lru.Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Get(1)
	obj.Put(3, 3)
	obj.Get(2)
	obj.Put(4, 4)
	obj.Get(1)
	obj.Get(3)
	obj.Get(4)
}
