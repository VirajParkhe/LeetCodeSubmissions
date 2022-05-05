import "sort"
func partitionLabels(s string) []int {
    m := map[string][]int{}
    for i:=0;i<len(s);i++{
        if val, ok := m[string(s[i])];!ok{
            m[string(s[i])] = []int{i,i}
        }else{
            val[1] = i
            m[string(s[i])] = val
        }
    }
    intervals := [][]int{}
    for _, v := range m {
        intervals = append(intervals, v)
    }
    sort.Slice(intervals, func(i,j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    parts := [][]int{}
    start :=  intervals[0][0]
    end := intervals[0][1]
    for i:=1;i<len(intervals);i++{
        if intervals[i][0] > end {
            parts = append(parts, []int{start, end})
            start = intervals[i][0]
            end = intervals[i][1]
        }else{
            if intervals[i][1] > end {
                end = intervals[i][1]
            }
        }
    }
    parts = append(parts, []int{start, end})
    ret := []int{}
    for i:=0;i<len(parts);i++{
        ret = append(ret, parts[i][1] - parts[i][0] +1)
    }
    return ret
}