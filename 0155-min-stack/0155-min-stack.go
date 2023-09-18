type MinStack struct {
    stack []struct {
        val, min int
    }
}

func Constructor() MinStack {
    return MinStack{}
}

func (ms *MinStack) Push(val int) {
    min := val
    if len(ms.stack) > 0 && val > ms.GetMin() {
        min = ms.GetMin()
    }
    ms.stack = append(ms.stack, struct{ val, min int }{val, min})
}

func (ms *MinStack) Pop() {
    if len(ms.stack) > 0 {
        ms.stack = ms.stack[:len(ms.stack)-1]
    }
}

func (ms *MinStack) Top() int {
    if len(ms.stack) > 0 {
        return ms.stack[len(ms.stack)-1].val
    }
    return 0 // Handle this case as needed
}

func (ms *MinStack) GetMin() int {
    if len(ms.stack) > 0 {
        return ms.stack[len(ms.stack)-1].min
    }
    return 0 // Handle this case as needed
}
