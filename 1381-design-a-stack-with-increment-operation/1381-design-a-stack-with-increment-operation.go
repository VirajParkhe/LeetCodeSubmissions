type CustomStack struct {
    stk []int
    l int
    max int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{[]int{},0,maxSize}
}


func (this *CustomStack) Push(x int)  {
    if this.l != this.max{
        this.stk = append(this.stk, x)
        this.l +=1
    }
}


func (this *CustomStack) Pop() int {
    if this.l!=0{
        val := this.stk[this.l-1]
        this.stk = this.stk[:this.l-1]
        this.l -=1
        return val
    }
    return -1
}


func (this *CustomStack) Increment(k int, val int)  {
    for i:=0;i<k && i<this.l;i++{
        this.stk[i] +=val 
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */