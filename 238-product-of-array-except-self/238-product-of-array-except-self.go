func productExceptSelf(nums []int) []int {
    pre := 1
    suf:= 1
    l := len(nums)
    ans := []int{}
    for i:=0;i<l;i++{
        ans = append(ans, 1)
    }
    for i:=0;i<l;i++{
        ans[i] *= pre
        pre *= nums[i]
        ans[l-i-1] *=suf
        suf  *= nums[l-i-1]
    }
    return ans
}