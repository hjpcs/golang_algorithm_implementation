package main

type MyCircularDeque struct {
	queue []int
	head  int
	tail  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: make([]int, k+1),
		head:  0,
		tail:  0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = (this.head - 1 + len(this.queue)) % len(this.queue)
	this.queue[this.head] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.tail] = value
	this.tail = (this.tail + 1) % len(this.queue)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.queue)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + len(this.queue)) % len(this.queue)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	n := (this.tail - 1 + len(this.queue)) % len(this.queue)
	return this.queue[n]
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if (this.tail+1)%len(this.queue) == this.head {
		return true
	}
	return false
}
