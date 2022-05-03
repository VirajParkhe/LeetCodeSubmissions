import "sort"
type sortedInterval [][]int

func (this sortedInterval) Less(i, j int)bool {
    return this[i][0] < this[j][0]
}

func (this sortedInterval) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]    
}

func (this sortedInterval) Len()int {
    return len(this)    
}

func maxWidthOfVerticalArea(points [][]int) int {
    temp := sortedInterval(points)
    sort.Sort(temp)
    max := 0
    for i:=0;i<len(temp)-1;i++{
        if  (temp[i+1][0] - temp[i][0] ) > max {
            max = temp[i+1][0] - temp[i][0]
        }
    }
    return max
}