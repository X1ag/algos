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
	// if l.Head != nil {
	// 	prev := l.Head.Next
	// 	l.Head = &Node{Key: k, Value: v, Next: prev}
	// 	return
	// }
	// prev := l.Head
	// l.Head = &Node{Key: k, Value: v, Next: prev}
}

type MyHashMap struct {
  buckets []*LinkedList
	capacity int
}

func (this *MyHashMap) getBucketIndex(key int) int {
	return key % this.capacity
}

func Constructor() MyHashMap {
	capacity := 2;
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
   return -1 
}


func (this *MyHashMap) Remove(key int)  {
    
}