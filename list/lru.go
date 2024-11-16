package list

import "container/list"

/**
146. LRU 缓存

请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

/** 两个点需要注意
1. map 的 key 是 put 的 key，value 是 list 的元素
2. list 的元素存的是 key 和 value，不能只存 value，否则删除（清理）的时候有麻烦
*/

type KV struct {
    Key int
    Val int
}

type LRUCache struct {
    cap   int
    cache map[int]*list.Element
    list  *list.List
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        cap:   capacity,
        cache: make(map[int]*list.Element),
        list:  list.New(),
    }
}

func (this *LRUCache) Get(key int) int {
    if ele, ok := this.cache[key]; ok {
        this.list.MoveToFront(ele)
        return ele.Value.(*KV).Val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if ele, ok := this.cache[key]; ok {
        ele.Value.(*KV).Val = value
        this.list.MoveToFront(ele)
    } else {
        // 注意这里是>=
        if len(this.cache) >= this.cap {
            back := this.list.Back()
            delete(this.cache, back.Value.(*KV).Key)
            this.list.Remove(back)
        }
        ele := this.list.PushFront(&KV{key, value})
        this.cache[key] = ele
    }
}
