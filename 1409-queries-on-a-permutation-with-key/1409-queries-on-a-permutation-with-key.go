func findIdx(n, l int, nums []int)int{
    for i:=0;i<l;i++{
        if nums[i] == n {
            return i
        }
    }
    return -1
}

func processQueries(queries []int, m int) []int {
    p := make([]int, m)
    for i:=0;i<m;i++{
        p[i] = i+1
    }
    ret := []int{}
    for i:=0;i<len(queries);i++{
        idx := findIdx(queries[i], m, p)
        ret = append(ret, idx)
        if idx != 0 {
            temp := append([]int{queries[i]}, p[:idx]...)
            if idx != m-1 {
                p = append(temp, p[idx+1:]...) 
            }else{
                p = temp
            }
        }
    }
    return ret
}