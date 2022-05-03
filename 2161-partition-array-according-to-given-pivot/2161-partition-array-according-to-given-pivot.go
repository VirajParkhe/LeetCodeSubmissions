func pivotArray(nums []int, pivot int) []int {
    l := []int{}
    m := []int{}
    r := []int{}
    for _, v := range nums {
        if v < pivot {
            l = append(l, v)
        }else if v > pivot {
            r = append(r, v)
        }else {
            m = append(m,v)
        }
    }
    return append(l, append(m, r...)...)
}