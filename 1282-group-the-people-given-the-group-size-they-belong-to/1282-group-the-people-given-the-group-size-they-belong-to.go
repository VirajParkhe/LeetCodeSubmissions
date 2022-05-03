func groupThePeople(groupSizes []int) [][]int {
    m := map[int][]int{}
    for i, size := range groupSizes {
        m[size] = append(m[size], i)
    }
    ret := [][]int{}
    for size, lst := range m {
        for i:=0;i<=len(lst)-size;i = i+size{
            ret = append(ret, lst[i:i+size])
        }
    }
    return ret
}