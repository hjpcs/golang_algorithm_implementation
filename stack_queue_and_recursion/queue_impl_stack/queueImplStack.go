package queue_impl_stack

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	head := this.Top()
	if head == -1 {
		return -1
	}
	this.queue = this.queue[:len(this.queue)-1]
	return head
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
