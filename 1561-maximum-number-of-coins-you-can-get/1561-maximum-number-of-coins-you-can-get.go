func maxCoins(piles []int) int {
    sort.Ints(piles)
    l := len(piles)
    j :=  l -2
    sum := 0
    for i:=0;i<l/3;i++{
        sum += piles[j]
        j = j-2
    }
    return sum
}