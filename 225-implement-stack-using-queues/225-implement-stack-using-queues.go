type MyStack struct {
    l []int
    sl int
}


func Constructor() MyStack {
    return MyStack{[]int{}, 0}
}


func (this *MyStack) Push(x int)  {
    this.l = append(this.l, x)
    this.sl += 1
}


func (this *MyStack) Pop() int {
    temp := 0
    if this.sl != 0 {
        temp = this.l[this.sl-1]
        this.l = this.l[:this.sl - 1]
        this.sl -=1
    }
    return temp
}


func (this *MyStack) Top() int {
    if this.sl != 0 {
        return this.l[this.sl-1]
    }
    return -1
}


func (this *MyStack) Empty() bool {
    return this.sl == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */