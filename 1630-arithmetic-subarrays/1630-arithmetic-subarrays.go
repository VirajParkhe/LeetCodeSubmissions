func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
    ret := []bool{}
    for i:=0;i<len(l);i++{
        temp := make([]int, r[i]-l[i]+1)
        copy(temp, nums[l[i]:r[i]+1])
        sort.Ints(temp)
        start := -1
        isArith := true
        for j:=1;j<len(temp);j++{
            if start == -1 {
                start = temp[j] - temp[j-1]
            }else if  (temp[j] - temp[j-1] ) != start {
                isArith=false
                break
            }
        }
        if isArith {
            ret = append(ret, true)
        }else {
            ret = append(ret, false)
        }
    }
    return ret
}