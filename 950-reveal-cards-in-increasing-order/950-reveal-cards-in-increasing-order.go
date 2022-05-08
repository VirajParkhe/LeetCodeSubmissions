import "sort"


func deckRevealedIncreasing(deck []int) []int {
    l := len(deck)
    sort.Ints(deck)
    ret := make([]int,l)
    ind := make([]int, l)
    for i:=0;i<l;i++{
        ind[i] = i
    }
    for _, card := range deck {
        top := ind[0]
        ind = ind[1:]
        ret[top] = card
        if len(ind) > 1 {
            top := ind[0]
            ind = ind[1:]
            ind = append(ind, top)
        }
    }
    
    return ret
}