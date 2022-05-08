/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    flattenedList []*NestedInteger
    ind int
    l int
}

func flatten(nList []*NestedInteger) []*NestedInteger {
    ret := []*NestedInteger{}
    for i:=0;i<len(nList);i++{
        if nList[i].IsInteger() {
            ret = append(ret, nList[i])
        }else{
            ret = append(ret, flatten(nList[i].GetList())...)
        }
    }
    return ret
}

func printLst(nList []*NestedInteger) {
    for i:=0;i<len(nList);i++{
        fmt.Printf("%d->", nList[i].GetInteger())
    }
    fmt.Println(" ")
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    lst := flatten(nestedList)
    return &NestedIterator{lst, 0, len(lst)}
}

func (this *NestedIterator) Next() int {
    val := this.flattenedList[this.ind].GetInteger()
    this.ind++
    return val
}

func (this *NestedIterator) HasNext() bool {
     if this.ind != this.l {
        return true
    }
    return false
}