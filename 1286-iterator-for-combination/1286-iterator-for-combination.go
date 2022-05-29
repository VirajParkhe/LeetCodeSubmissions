func generateCombination(comb string, s []string){
    
}

type CombinationIterator struct {
    l int
    comb []string
    char string
    charLen int
    start int
    combLen int
}


func Constructor(characters string, combinationLength int) CombinationIterator {
    x:=CombinationIterator{combinationLength, []string{},characters,len(characters), 0, 0}
    x.generate("", 0, 0)
    return x
}


func (this *CombinationIterator) Next() string {
    ret := this.comb[this.start]
    this.start++
    return ret
}


func (this *CombinationIterator) HasNext() bool {
    if this.start < this.combLen{
        return true
    }
    return false
}

func (this *CombinationIterator) generate(s string, start, n int) {
    if n == this.l {
        if len(s) == this.l{
            this.comb = append(this.comb, s)
            this.combLen++   
        }
        return
    }
    for i:=start;i<this.charLen;i++{
        this.generate(s+string(this.char[i]), i+1,n+1)
    }
}


/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */