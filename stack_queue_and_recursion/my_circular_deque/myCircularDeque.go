package main

type MyCircularDeque struct {
	Queue []int
	N     int
	Head  int
	Tail  int
}

func Constructor(k int) MyCircularDeque {
	var queue MyCircularDeque
	var s []int
	queue.Queue = s
	queue.N = k
	return queue
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.N == len(this.Queue) {
		return false
	}
	if this.Head == 0 && this.Tail == 0 {
		this.Queue = append(this.Queue, value)
		this.Tail = 1
		return true
	}
	if this.Head > 0 {
		this.Head -= 1
		this.Queue[this.Head] = value
	} else {
		this.Head = this.N - 1
		this.Queue[this.Head] = value
	}
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.N == len(this.Queue) {
		return false
	}
	if this.Head == 0 && this.Tail == 0 {
		this.Queue = append(this.Queue, value)

		this.Tail = 1
		return true
	}
	this.Queue[this.Tail] = value
	this.Tail = (this.Tail + 1) % this.N
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.Queue) == 0 {
		return false
	}
	this.Head = (this.Head + 1) % this.N
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.Queue) == 0 {
		return false
	}
	if this.Tail > 0 {
		this.Tail -= 1
	} else {
		this.Tail = this.N - 1
	}
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if len(this.Queue) == 0 {
		return -1
	}
	return this.Queue[this.Head]
}

func (this *MyCircularDeque) GetRear() int {
	if len(this.Queue) == 0 {
		return -1
	}
	if this.Tail > 0 {
		return this.Queue[this.Tail-1]
	} else {
		return this.Queue[this.N-1]
	}
}

func (this *MyCircularDeque) IsEmpty() bool {
	if len(this.Queue) == 0 {
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if len(this.Queue) == this.N {
		return true
	}
	return false
}
