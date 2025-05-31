package main

type Node struct {
	Key int
	Value int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) putAtHead(k, v int) {
	l.Head = &Node{Key: k, Value: v, Next: l.Head}
}

type MyHashMap struct {
    buckets []*LinkedList
    capacity int
}

func (this *MyHashMap) getBucketIndex(key int) int {
	return key % this.capacity
}

func Constructor() MyHashMap {
	capacity := 1009;
	buckets := make([]*LinkedList, capacity)
    for i := range buckets {
        buckets[i] = &LinkedList{}
    }	
  return MyHashMap{buckets: buckets, capacity: capacity} 
}


func (this *MyHashMap) Put(key int, value int)  {
	indexBucket := this.getBucketIndex(key)
	bucket := this.buckets[indexBucket]
	
	for node := bucket.Head; node != nil; {
		if key == node.Key {
			node.Value = value
			return
		}
		node = node.Next
	}
	bucket.putAtHead(key,value)
}


func (this *MyHashMap) Get(key int) int {
  indexBucket := this.getBucketIndex(key)
	bucket := this.buckets[indexBucket] 

	for node := bucket.Head; node != nil; {
		if key == node.Key {
			return node.Value
		}
		node = node.Next
	}
	return -1
}


func (this *MyHashMap) Remove(key int)  {
  indexBucket := this.getBucketIndex(key)
	bucket := this.buckets[indexBucket] 

	var prev *Node
	for node := bucket.Head; node != nil; node = node.Next {
		if key == node.Key {
			if prev == nil {
				bucket.Head = node.Next
				return
			}
			prev.Next = node.Next	
			return
		}
		prev = node
	}
}