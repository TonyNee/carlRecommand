package lru

import (
	"container/list"
	"fmt"
)

type entry struct {
	key, val int
}

type Cache struct {
	_capacity int
	_list     *list.List
	umap      map[int]*list.Element
}

func Constructor(capacity int) Cache {
	return Cache{capacity, list.New(), map[int]*list.Element{}}
}

func (c *Cache) Get(key int) int {
	tmp := c.umap[key]
	if tmp == nil {
		fmt.Printf("[R] 找不到键值为%d的节点\n", key)
		return -1
	}
	c._list.MoveToFront(tmp)
	fmt.Printf("[R] 找到键值为%d的节点，其值为%d\n", key, tmp.Value.(entry).val)
	return tmp.Value.(entry).val
}

func (c *Cache) Put(key int, value int) {
	tmp := c.umap[key]
	if tmp != nil {
		tmp.Value = entry{key, value}
		c._list.MoveToFront(tmp)
		fmt.Printf("[U] 更新键值为%d的节点，更新为%d\n", key, value)
		return
	}
	c.umap[key] = c._list.PushFront(entry{key, value})
	fmt.Printf("[C] 添加键值为%d的节点，其值为%d\n", key, value)
	if len(c.umap) > c._capacity {
		last := c._list.Remove(c._list.Back())
		delete(c.umap, last.(entry).key)
		fmt.Printf("[D] 删除键值为%d的节点，其值为%d\n", last.(entry).key, last.(entry).val)
	}
}
