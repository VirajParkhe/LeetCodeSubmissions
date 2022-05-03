func rearrangeArray(nums []int) []int {
    n := []int{}
    p := []int{}
    for i:=0;i<len(nums);i++{
        if nums[i] < 0 {
            n = append(n, nums[i])
        }else {
            p= append(p, nums[i])
        }
    }
    ret := []int{}
    for i:=0;i<len(p);i++{
        ret = append(ret, []int{p[i], n[i]}...)
    }
    return ret
}