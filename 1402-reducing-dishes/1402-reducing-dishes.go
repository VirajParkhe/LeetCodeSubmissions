import "sort"
func maxSatisfaction(satisfaction []int) int {
    sort.Slice(satisfaction, func(i, j int)bool{
        return satisfaction[i] > satisfaction[j]
    })
    var i int
    sum := 0
    for i=0;i<len(satisfaction);i++{
        if satisfaction[i] <=0 {
            break
        }else{
            sum +=satisfaction[i]
        }
    }
    if i-1 < 0 {
        return 0
    } 
    prod := 0
    for j,count:=0, i;j<i;j,count = j+1,count-1{
        prod += count*satisfaction[j]
    }
    max := prod
    for j:=i;j<len(satisfaction);j++{
        prod = prod + sum + satisfaction[j]
        if prod > max {
            max = prod
        }
        sum += satisfaction[j]
    }
    return max
}