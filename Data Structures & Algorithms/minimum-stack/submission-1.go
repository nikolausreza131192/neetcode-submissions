type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt32
	for i := 0; i < len(this.stack); i++ {
		if this.stack[i] < min {
			min = this.stack[i]
		}
	}
	return min
}
