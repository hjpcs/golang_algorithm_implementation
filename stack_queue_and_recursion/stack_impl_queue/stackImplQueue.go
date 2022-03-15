package stack_impl_queue

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if this.Peek() != -1 {
		head := this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
		return head
	}
	return -1
}

func (this *MyQueue) Peek() int {
	if !this.Empty() {
		if len(this.outStack) != 0 {
			return this.outStack[len(this.outStack)-1]
		} else {
			for len(this.inStack) > 0 {
				this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
				this.inStack = this.inStack[:len(this.inStack)-1]
			}
			return this.outStack[len(this.outStack)-1]
		}
	}
	return -1
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}
