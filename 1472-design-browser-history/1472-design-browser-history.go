type BrowserHistory struct {
    hist []string
    curr int
    l  int
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{[]string{homepage}, 0, 1}
}


func (this *BrowserHistory) Visit(url string)  {
    if this.curr != this.l-1 && this.l!=0 {
        this.hist = this.hist[:this.curr+1]
        this.l = this.curr+1
    }
    this.hist = append(this.hist, url)
    this.curr++
    this.l++
}


func (this *BrowserHistory) Back(steps int) string {
    if this.curr - steps >= 0 {
        this.curr -=steps
        return this.hist[this.curr]
    }
    this.curr =0 
    return this.hist[this.curr]
}


func (this *BrowserHistory) Forward(steps int) string {
    this.curr += steps
    if this.curr >= this.l {
        this.curr = this.l -1
    }
    return this.hist[this.curr]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */